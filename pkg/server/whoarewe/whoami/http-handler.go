package whoami

import (
	"encoding/json"
	"fmt"
	"github.com/watergist/k8s-manifests/pkg/server"
	"github.com/watergist/k8s-manifests/pkg/server/listener"
	"github.com/watergist/k8s-manifests/pkg/server/whoarewe/whoami/wide"
	"log"
	"net/http"
)

type Server struct {
	*listener.Server
}

func (s *Server) RequestHost(w http.ResponseWriter, r *http.Request) {
	// https://stackoverflow.com/a/51477680
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := w.Write([]byte(fmt.Sprintf("You are at \"%v\"\n", r.Host)))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func (s *Server) RequestIP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("Your address \"%v\"\n", r.RemoteAddr)))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func (s *Server) RequestInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	requestInfo := wide.GenRequestInfo(r)
	h := w.Header()
	requestInfo.BasicResponseHeaders = &h
	requestInfoBytes, err := json.MarshalIndent(requestInfo, "", "  ")
	if err != nil {
		server.WriteServerError(w, err, "reading labels", http.StatusNotImplemented)
		return
	}
	_, err = w.Write(requestInfoBytes)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
