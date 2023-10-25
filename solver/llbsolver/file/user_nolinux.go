//go:build !linux
// +build !linux

package file

import (
	"github.com/pkg/errors"
	"github.com/preminger/buildkit/solver/llbsolver/ops/fileoptypes"
	"github.com/preminger/buildkit/solver/pb"
	copy "github.com/tonistiigi/fsutil/copy"
)

func readUser(chopt *pb.ChownOpt, mu, mg fileoptypes.Mount) (*copy.User, error) {
	if chopt == nil {
		return nil, nil
	}
	return nil, errors.New("only implemented in linux")
}
