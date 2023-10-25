package client

import (
	"context"

	"github.com/pkg/errors"
	controlapi "github.com/preminger/buildkit/api/services/control"
	apitypes "github.com/preminger/buildkit/api/types"
)

type Info struct {
	BuildkitVersion BuildkitVersion `json:"buildkitVersion"`
}

type BuildkitVersion struct {
	Package  string `json:"package"`
	Version  string `json:"version"`
	Revision string `json:"revision"`
}

func (c *Client) Info(ctx context.Context) (*Info, error) {
	res, err := c.ControlClient().Info(ctx, &controlapi.InfoRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to call info")
	}
	return &Info{
		BuildkitVersion: fromAPIBuildkitVersion(res.BuildkitVersion),
	}, nil
}

func fromAPIBuildkitVersion(in *apitypes.BuildkitVersion) BuildkitVersion {
	if in == nil {
		return BuildkitVersion{}
	}
	return BuildkitVersion{
		Package:  in.Package,
		Version:  in.Version,
		Revision: in.Revision,
	}
}
