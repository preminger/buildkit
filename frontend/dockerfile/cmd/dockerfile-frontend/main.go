package main

import (
	"flag"
	"fmt"
	"os"

	dockerfile "github.com/preminger/buildkit/frontend/dockerfile/builder"
	"github.com/preminger/buildkit/frontend/gateway/grpcclient"
	"github.com/preminger/buildkit/util/appcontext"
	"github.com/preminger/buildkit/util/stack"
	"github.com/sirupsen/logrus"
)

func init() {
	stack.SetVersionInfo(Version, Revision)
}

func main() {
	var version bool
	flag.BoolVar(&version, "version", false, "show version")
	flag.Parse()

	if version {
		fmt.Printf("%s %s %s %s\n", os.Args[0], Package, Version, Revision)
		os.Exit(0)
	}

	if err := grpcclient.RunFromEnvironment(appcontext.Context(), dockerfile.Build); err != nil {
		logrus.Errorf("fatal error: %+v", err)
		panic(err)
	}
}
