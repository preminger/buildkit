package dockerfile2llb

import (
	ocispecs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/preminger/buildkit/exporter/containerimage/image"
	"github.com/preminger/buildkit/util/system"
)

// Image is the JSON structure which describes some basic information about the image.
// This provides the `application/vnd.oci.image.config.v1+json` mediatype when marshalled to JSON.
type Image image.Image

func clone(src Image) Image {
	img := src
	img.Config = src.Config
	img.Config.Env = append([]string{}, src.Config.Env...)
	img.Config.Cmd = append([]string{}, src.Config.Cmd...)
	img.Config.Entrypoint = append([]string{}, src.Config.Entrypoint...)
	return img
}

func emptyImage(platform ocispecs.Platform) Image {
	img := Image{}
	img.Platform.Architecture = platform.Architecture
	img.Platform.OS = platform.OS
	img.Platform.Variant = platform.Variant
	img.RootFS.Type = "layers"
	img.Config.WorkingDir = "/"
	img.Config.Env = []string{"PATH=" + system.DefaultPathEnv(platform.OS)}
	return img
}
