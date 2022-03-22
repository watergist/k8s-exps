//go:build !local

package main

import (
	_ "embed"
	multiserver "github.com/watergist/k8s-manifests/pkg/whoarewe/multi-server"
)

func registerLocalCertificate(s *multiserver.Server) {
	// pass
}
