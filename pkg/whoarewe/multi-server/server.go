package multiserver

import (
	"context"
	"crypto/tls"
	"github.com/spf13/viper"
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
	TLSPath     string
	Certificate *tls.Certificate
	HTTPServer  map[string]*http.Server
	HTTPSServer map[string]*http.Server
	WG          sync.WaitGroup
}

func GetServer() *Server {
	s := Server{
		HTTPServer:  map[string]*http.Server{},
		HTTPSServer: map[string]*http.Server{},
		TLSPath:     viper.GetString("TLS_KEYPAIR_PATH"),
	}
	for _, p := range strings.Split(viper.GetString("HTTP_PORTS"), "~") {
		p = strings.TrimSpace(p)
		if p != "" {
			s.HTTPServer[p] = s.createServer(p, "HTTP")
		}
	}
	for _, p := range strings.Split(viper.GetString("HTTPS_PORTS"), "~") {
		p = strings.TrimSpace(p)
		if p != "" {
			s.HTTPSServer[p] = s.createServer(p, "HTTPS")
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
	return &http.Server{Addr: "0.0.0.0:" + port, Handler: regsiter.EnableDualLogging(mux, "0.0.0.0:"+port)}
}

func (s *Server) RunServers() {
	for p := range s.HTTPServer {
		s.WG.Add(1)
		go s.runHTTPServer(p)
	}
	for p := range s.HTTPSServer {
		s.WG.Add(1)
		go s.runHTTPSServer(p)
	}
}

func (s *Server) runHTTPServer(port string) {
	log.Printf("Listening HTTP on %v \n", port)
	err := s.HTTPServer[port].ListenAndServe()
	if err != nil {
		s.WG.Done()
		log.Fatalf("Error serving on this HTTP port %v: %v", port, err.Error())
		return
	}
}

func (s *Server) runHTTPSServer(port string) {
	log.Printf("Listening HTTPS on %v \n", port)
	var tlsKeyPath, tlsCertPath string
	if s.Certificate == nil {
		tlsCertPath = path.Join(s.TLSPath, certificateName)
		tlsKeyPath = path.Join(s.TLSPath, privateKeyName)
	} else {
		cfg := &tls.Config{Certificates: []tls.Certificate{*s.Certificate}}
		s.HTTPSServer[port].TLSConfig = cfg
	}
	err := s.HTTPSServer[port].ListenAndServeTLS(tlsCertPath, tlsKeyPath)
	if err != nil {
		s.WG.Done()
		log.Fatalf("Error serving on this HTTPS port %v: %v", port, err.Error())
		return
	}
}

func (s *Server) StopServer(port string) error {
	log.Printf("ShuttingDown server on %v \n", port)
	if _, ok := s.HTTPServer[port]; ok {
		err := s.HTTPServer[port].Shutdown(context.TODO())
		if err != nil {
			log.Fatalf("Error shuttingDown server on this port %v: %v", port, err.Error())
			return err
		}
		s.WG.Done()
	}
	if _, ok := s.HTTPSServer[port]; ok {
		err := s.HTTPSServer[port].Shutdown(context.TODO())
		if err != nil {
			log.Fatalf("Error shuttingDown server on this port %v: %v", port, err.Error())
			return err
		}
		s.WG.Done()
	}
	return nil
}
