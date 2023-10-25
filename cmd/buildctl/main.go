package main

import (
	"fmt"
	"os"

	_ "github.com/preminger/buildkit/client/connhelper/dockercontainer"
	_ "github.com/preminger/buildkit/client/connhelper/kubepod"
	_ "github.com/preminger/buildkit/client/connhelper/podmancontainer"
	_ "github.com/preminger/buildkit/client/connhelper/ssh"
	bccommon "github.com/preminger/buildkit/cmd/buildctl/common"
	"github.com/preminger/buildkit/solver/errdefs"
	"github.com/preminger/buildkit/util/apicaps"
	"github.com/preminger/buildkit/util/appdefaults"
	"github.com/preminger/buildkit/util/profiler"
	"github.com/preminger/buildkit/util/stack"
	_ "github.com/preminger/buildkit/util/tracing/detect/delegated"
	_ "github.com/preminger/buildkit/util/tracing/detect/jaeger"
	_ "github.com/preminger/buildkit/util/tracing/env"
	"github.com/preminger/buildkit/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"go.opentelemetry.io/otel"
)

func init() {
	apicaps.ExportedProduct = "buildkit"

	stack.SetVersionInfo(version.Version, version.Revision)

	// do not log tracing errors to stdio
	otel.SetErrorHandler(skipErrors{})
}

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Name, version.Package, c.App.Version, version.Revision)
	}
	app := cli.NewApp()
	app.Name = "buildctl"
	app.Usage = "build utility"
	app.Version = version.Version

	defaultAddress := os.Getenv("BUILDKIT_HOST")
	if defaultAddress == "" {
		defaultAddress = appdefaults.Address
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output in logs",
		},
		cli.StringFlag{
			Name:  "addr",
			Usage: "buildkitd address",
			Value: defaultAddress,
		},
		cli.StringFlag{
			Name:  "tlsservername",
			Usage: "buildkitd server name for certificate validation",
			Value: "",
		},
		cli.StringFlag{
			Name:  "tlscacert",
			Usage: "CA certificate for validation",
			Value: "",
		},
		cli.StringFlag{
			Name:  "tlscert",
			Usage: "client certificate",
			Value: "",
		},
		cli.StringFlag{
			Name:  "tlskey",
			Usage: "client key",
			Value: "",
		},
		cli.StringFlag{
			Name:  "tlsdir",
			Usage: "directory containing CA certificate, client certificate, and client key",
			Value: "",
		},
		cli.IntFlag{
			Name:  "timeout",
			Usage: "timeout backend connection after value seconds",
			Value: 5,
		},
	}

	app.Commands = []cli.Command{
		diskUsageCommand,
		pruneCommand,
		buildCommand,
		debugCommand,
		dialStdioCommand,
	}

	var debugEnabled bool

	app.Before = func(context *cli.Context) error {
		debugEnabled = context.GlobalBool("debug")

		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
		if debugEnabled {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	if err := bccommon.AttachAppContext(app); err != nil {
		handleErr(debugEnabled, err)
	}

	profiler.Attach(app)

	handleErr(debugEnabled, app.Run(os.Args))
}

func handleErr(debug bool, err error) {
	if err == nil {
		return
	}
	for _, s := range errdefs.Sources(err) {
		s.Print(os.Stderr)
	}
	if debug {
		fmt.Fprintf(os.Stderr, "error: %+v", stack.Formatter(err))
	} else {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	os.Exit(1)
}

type skipErrors struct{}

func (skipErrors) Handle(err error) {}
