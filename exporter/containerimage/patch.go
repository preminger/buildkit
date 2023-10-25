//go:build !nydus
// +build !nydus

package containerimage

import (
	"context"

	ocispecs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/preminger/buildkit/cache"
	"github.com/preminger/buildkit/session"
	"github.com/preminger/buildkit/solver"
)

func patchImageLayers(ctx context.Context, remote *solver.Remote, history []ocispecs.History, ref cache.ImmutableRef, opts *ImageCommitOpts, sg session.Group) (*solver.Remote, []ocispecs.History, error) {
	remote, history = normalizeLayersAndHistory(ctx, remote, history, ref, opts.OCITypes)
	return remote, history, nil
}
