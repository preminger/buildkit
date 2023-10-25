//go:build windows
// +build windows

package main

import (
	"crypto/tls"
	"net"

	"github.com/pkg/errors"
	_ "github.com/preminger/buildkit/solver/llbsolver/ops"
)

func listenFD(addr string, tlsConfig *tls.Config) (net.Listener, error) {
	return nil, errors.New("listening server on fd not supported on windows")
}
