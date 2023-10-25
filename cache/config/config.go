package config

import "github.com/preminger/buildkit/util/compression"

type RefConfig struct {
	Compression            compression.Config
	PreferNonDistributable bool
}
