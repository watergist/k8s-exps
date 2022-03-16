package multiserver

import (
	"context"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/listener"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/regsiter"
	"log"
	"net/http"
	"path"
	"strings"
	"sync"
)

const (
	privateKeyName  = "tls.key"
	certificateName = "tls.crt"
)

type Server struct {
	HttpPorts   []string
	HttpsPorts  []string
	TLSPath     string
	HttpServer  map[string]*http.Server
	HttpsServer map[string]*http.Server
	WG          sync.WaitGroup
}

func GetServer(httpPorts, httpsPorts string) *Server {
	s := Server{
		HttpPorts:   strings.Split(httpPorts, "~"),
		HttpsPorts:  strings.Split(httpsPorts, "~"),
		HttpServer:  map[string]*http.Server{},
		HttpsServer: map[string]*http.Server{},
	}
	for _, p := range s.HttpPorts {
		p = strings.TrimSpace(p)
		if p != "" {
			s.HttpServer[p] = s.createServer(p, "HTTP")
		}
	}
	for _, p := range s.HttpsPorts {
		p = strings.TrimSpace(p)
		if p != "" {
			s.HttpsServer[p] = s.createServer(p, "HTTPS")
		}
	}
	return &s
}

func (s *Server) createServer(port string, protocol string) *http.Server {
	mux := http.NewServeMux()
	regsiter.RegisterEndpoints(mux, &listener.Listener{
		Address:  "0.0.0.0",
		Port:     port,
		Protocol: protocol,
	})
	return &http.Server{Addr: "0.0.0.0:" + port, Handler: mux}
}

func (s *Server) RunServers() {
	for _, p := range s.HttpPorts {
		s.WG.Add(1)
		go s.runHttpServer(p)
	}
	for _, p := range s.HttpsPorts {
		s.WG.Add(1)
		go s.runHttpsServer(p)
	}
}

func (s *Server) runHttpServer(port string) {
	log.Printf("Listening HTTP on %v \n", port)
	err := s.HttpServer[port].ListenAndServe()
	if err != nil {
		s.WG.Done()
		log.Fatalf("Error serving on this HTTP port %v: %v", port, err.Error())
		return
	}
}

func (s *Server) runHttpsServer(port string) {
	log.Printf("Listening HTTPS on %v \n", port)
	err := s.HttpsServer[port].ListenAndServeTLS(path.Join(s.TLSPath, certificateName), path.Join(s.TLSPath, privateKeyName))
	if err != nil {
		s.WG.Done()
		log.Fatalf("Error serving on this HTTPS port %v: %v", port, err.Error())
		return
	}
}

func (s *Server) StopServer(port string) error {
	log.Printf("ShuttingDown server on %v \n", port)
	if _, ok := s.HttpServer[port]; ok {
		err := s.HttpServer[port].Shutdown(context.TODO())
		if err != nil {
			log.Fatalf("Error shuttingDown server on this port %v: %v", port, err.Error())
			return err
		}
		s.WG.Done()
	}
	if _, ok := s.HttpsServer[port]; ok {
		err := s.HttpsServer[port].Shutdown(context.TODO())
		if err != nil {
			log.Fatalf("Error shuttingDown server on this port %v: %v", port, err.Error())
			return err
		}
		s.WG.Done()
	}
	return nil
}
