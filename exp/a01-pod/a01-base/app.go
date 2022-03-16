package main

import (
	"github.com/watergist/k8s-manifests/pkg/whoarewe/listener"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/regsiter"
	"log"
	"net/http"
)

const PORT = "3001"

func main() {
	log.Println("Started Application")
	mux := http.NewServeMux()
	regsiter.RegisterEndpoints(mux, &listener.Listener{
		Address:  "0.0.0.0",
		Port:     PORT,
		Protocol: "HTTP",
	})
	if err := http.ListenAndServe("0.0.0.0:"+PORT, mux); err != nil {
		log.Fatal(err)
	}
	log.Println("Exited Application")
}
