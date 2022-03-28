package whoareu

import (
	"encoding/json"
	"fmt"
	"github.com/watergist/k8s-manifests/pkg/server"
	"github.com/watergist/k8s-manifests/pkg/server/listener"
	"github.com/watergist/k8s-manifests/pkg/server/whoarewe/whoareu/wide"
	"log"
	"net/http"
	"os"
)

type Server struct {
	*listener.Server
}

func (s *Server) PodName(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf(
		"\"%v\" Serving at %v:%v from \"%v\"\n",
		os.Getenv("POD_NAME"), s.ListenerProperties.Address, s.ListenerProperties.Port, os.Getenv("POD_NAMESPACE"))))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func (s *Server) PodInfo(w http.ResponseWriter, _ *http.Request) {
	podProperties := wide.GetPodProperties(s.ListenerProperties)
	podPropertiesBytes, err := json.MarshalIndent(podProperties, "", "  ")
	if err != nil {
		server.WriteServerError(w, err, "reading labels", http.StatusNotImplemented)
		return
	}
	w.Header().Add("content-type", "application/json")
	_, err = w.Write(podPropertiesBytes)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
