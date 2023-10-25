package forwarder

import (
	"context"

	"github.com/preminger/buildkit/frontend"
	"github.com/preminger/buildkit/frontend/gateway/client"
	"github.com/preminger/buildkit/session"
	"github.com/preminger/buildkit/solver/pb"
	"github.com/preminger/buildkit/worker"
)

func NewGatewayForwarder(w worker.Infos, f client.BuildFunc) frontend.Frontend {
	return &GatewayForwarder{
		workers: w,
		f:       f,
	}
}

type GatewayForwarder struct {
	workers worker.Infos
	f       client.BuildFunc
}

func (gf *GatewayForwarder) Solve(ctx context.Context, llbBridge frontend.FrontendLLBBridge, opts map[string]string, inputs map[string]*pb.Definition, sid string, sm *session.Manager) (retRes *frontend.Result, retErr error) {
	c, err := llbBridgeToGatewayClient(ctx, llbBridge, opts, inputs, gf.workers, sid, sm)
	if err != nil {
		return nil, err
	}

	defer func() {
		c.discard(retErr)
	}()

	res, err := gf.f(ctx, c)
	if err != nil {
		return nil, err
	}

	return c.toFrontendResult(res)
}
