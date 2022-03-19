package whoareu

import (
	"encoding/json"
	"fmt"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/listener"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/whoareu/wide"
	"log"
	"net/http"
	"os"
)

type Server struct {
	ListenerProperties *listener.Listener
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
		writeServerError(w, err, "reading labels", http.StatusNotImplemented)
		return
	}
	w.Header().Add("content-type", "application/json")
	_, err = w.Write(podPropertiesBytes)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func writeServerError(w http.ResponseWriter, err error, while string, status int) {
	w.WriteHeader(status)
	_, err = w.Write([]byte(fmt.Sprintf(
		"Error while %v: %v", while, err.Error(),
	)))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
