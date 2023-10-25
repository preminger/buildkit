package testutil

import (
	"testing"

	"github.com/preminger/buildkit/solver"
)

func TestMemoryCacheStorage(t *testing.T) {
	RunCacheStorageTests(t, solver.NewInMemoryCacheStorage)
}
