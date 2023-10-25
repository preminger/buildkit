package frontend

import (
	"context"

	digest "github.com/opencontainers/go-digest"
	"github.com/preminger/buildkit/client/llb"
	gw "github.com/preminger/buildkit/frontend/gateway/client"
	"github.com/preminger/buildkit/session"
	"github.com/preminger/buildkit/solver"
	"github.com/preminger/buildkit/solver/pb"
	"github.com/preminger/buildkit/solver/result"
)

type Result = result.Result[solver.ResultProxy]

type Attestation = result.Attestation[solver.ResultProxy]

type Frontend interface {
	Solve(ctx context.Context, llb FrontendLLBBridge, opt map[string]string, inputs map[string]*pb.Definition, sid string, sm *session.Manager) (*Result, error)
}

type FrontendLLBBridge interface {
	Solve(ctx context.Context, req SolveRequest, sid string) (*Result, error)
	ResolveImageConfig(ctx context.Context, ref string, opt llb.ResolveImageConfigOpt) (digest.Digest, []byte, error)
	Warn(ctx context.Context, dgst digest.Digest, msg string, opts WarnOpts) error
}

type SolveRequest = gw.SolveRequest

type CacheOptionsEntry = gw.CacheOptionsEntry

type WarnOpts = gw.WarnOpts
