package main

import (
	multiserver "github.com/watergist/k8s-manifests/pkg/whoarewe/multi-server"
	"log"
	"os"
)

func main() {
	log.Println("Started Application")
	s := multiserver.GetServer(os.Getenv("HTTP_PORTS"), os.Getenv("HTTPS_PORTS"))
	s.TLSPath = TLSPath
	s.RunServers()
	s.WG.Wait()
	log.Println("Exited Application")
}
