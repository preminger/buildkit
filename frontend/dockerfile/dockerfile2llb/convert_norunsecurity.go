//go:build !dfrunsecurity
// +build !dfrunsecurity

package dockerfile2llb

import (
	"github.com/preminger/buildkit/client/llb"
	"github.com/preminger/buildkit/frontend/dockerfile/instructions"
)

func dispatchRunSecurity(c *instructions.RunCommand) (llb.RunOption, error) {
	return nil, nil
}
