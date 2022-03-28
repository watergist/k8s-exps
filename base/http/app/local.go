package main

import (
	"crypto/tls"
	_ "embed"
	"github.com/watergist/k8s-manifests/pkg/server/multi-server"
	"log"
)

//go:embed testdata/tls.key
var tlsKey []byte

//go:embed testdata/tls.crt
var tlsCrt []byte

func registerLocalCertificate(s *multiserver.Server) {
	certificate, err := tls.X509KeyPair(tlsCrt, tlsKey)
	if err != nil {
		log.Fatalf("Error generating certificate data from compile time saved keyPair: %v\n", err)
	}
	s.Certificate = &certificate
}
