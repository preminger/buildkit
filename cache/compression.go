//go:build !nydus
// +build !nydus

package cache

import (
	"context"

	"github.com/containerd/containerd/content"
	ocispecs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/preminger/buildkit/cache/config"
)

func needsForceCompression(ctx context.Context, cs content.Store, source ocispecs.Descriptor, refCfg config.RefConfig) bool {
	return refCfg.Compression.Force
}
