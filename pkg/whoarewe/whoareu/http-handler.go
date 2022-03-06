package whoareu

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"watergist/k8s-manifests/pkg/whoarewe/whoareu/wide"
)

func PodName(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("I am \"%v\" from \"%v\"\n", os.Getenv("POD_NAME"), os.Getenv("POD_NAMESPACE"))))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func PodInfo(w http.ResponseWriter, _ *http.Request) {
	podProperties := wide.GetPodProperties()
	podPropertiesBytes, err := json.Marshal(podProperties)
	if err != nil {
		writeServerError(w, err, "reading labels", http.StatusNotImplemented)
		return
	}
	w.Header().Add("a-forwArDed-by", "whoarewe")
	w.Header().Add("a-forwArDed-by", os.Getenv("POD_NAME"))
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
