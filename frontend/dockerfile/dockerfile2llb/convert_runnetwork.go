package dockerfile2llb

import (
	"github.com/pkg/errors"

	"github.com/preminger/buildkit/client/llb"
	"github.com/preminger/buildkit/frontend/dockerfile/instructions"
	"github.com/preminger/buildkit/solver/pb"
)

func dispatchRunNetwork(c *instructions.RunCommand) (llb.RunOption, error) {
	network := instructions.GetNetwork(c)

	switch network {
	case instructions.NetworkDefault:
		return nil, nil
	case instructions.NetworkNone:
		return llb.Network(pb.NetMode_NONE), nil
	case instructions.NetworkHost:
		return llb.Network(pb.NetMode_HOST), nil
	default:
		return nil, errors.Errorf("unsupported network mode %q", network)
	}
}
