package llbsolver

import (
	"context"

	"github.com/pkg/errors"
	cacheconfig "github.com/preminger/buildkit/cache/config"
	"github.com/preminger/buildkit/frontend"
	"github.com/preminger/buildkit/session"
	"github.com/preminger/buildkit/solver"
	"github.com/preminger/buildkit/solver/llbsolver/provenance"
	"github.com/preminger/buildkit/worker"
)

type Result struct {
	*frontend.Result
	Provenance *provenance.Result
}

type Attestation = frontend.Attestation

func workerRefResolver(refCfg cacheconfig.RefConfig, all bool, g session.Group) func(ctx context.Context, res solver.Result) ([]*solver.Remote, error) {
	return func(ctx context.Context, res solver.Result) ([]*solver.Remote, error) {
		ref, ok := res.Sys().(*worker.WorkerRef)
		if !ok {
			return nil, errors.Errorf("invalid result: %T", res.Sys())
		}

		return ref.GetRemotes(ctx, true, refCfg, all, g)
	}
}
