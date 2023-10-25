package imageutil

import (
	"encoding/base64"
	"encoding/json"

	"github.com/pkg/errors"
	binfotypes "github.com/preminger/buildkit/util/buildinfo/types"
)

// BuildInfo returns build info from image config.
//
// Deprecated: Build information is deprecated: https://github.com/preminger/buildkit/blob/master/docs/deprecated.md
func BuildInfo(dt []byte) (*binfotypes.BuildInfo, error) {
	if len(dt) == 0 {
		return nil, nil
	}
	var config binfotypes.ImageConfig
	if err := json.Unmarshal(dt, &config); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal image config")
	}
	if len(config.BuildInfo) == 0 {
		return nil, nil
	}
	dtbi, err := base64.StdEncoding.DecodeString(config.BuildInfo)
	if err != nil {
		return nil, err
	}
	var bi binfotypes.BuildInfo
	if err = json.Unmarshal(dtbi, &bi); err != nil {
		return nil, errors.Wrap(err, "failed to decode buildinfo from image config")
	}
	return &bi, nil
}
