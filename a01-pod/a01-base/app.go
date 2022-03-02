package main

import (
	"log"
	"net/http"
	"watergist/k8s-manifests/pkg/whoami"
)

const PORT = "3001"

func main() {
	log.Println("Started Application")
	mux := http.NewServeMux()
	whoami.RegisterEndpoints(mux)
	if err := http.ListenAndServe("0.0.0.0:"+PORT, mux); err != nil {
		log.Fatal(err)
	}
	log.Println("Exited Application")
}
