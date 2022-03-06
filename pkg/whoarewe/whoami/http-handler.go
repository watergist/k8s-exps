package whoami

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"watergist/k8s-manifests/pkg/whoarewe/whoami/wide"
)

func RequestHost(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("You are at \"%v\" and wanted to be at \"%v\"\n", r.Host, r.URL.Host)))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func RequestIP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("Your address '\"%v\"\n", r.RemoteAddr)))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func RequestInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("a-forwArDed-by", "whoarewe")
	w.Header().Add("a-forwArDed-by", os.Getenv("POD_NAME"))
	w.Header().Add("content-type", "application/json")
	requestInfo := wide.GenRequestInfo(r)
	h := w.Header()
	requestInfo.BasicResponseHeaders = &h
	requestInfoBytes, err := json.Marshal(requestInfo)
	if err != nil {
		writeServerError(w, err, "reading labels", http.StatusNotImplemented)
		return
	}
	_, err = w.Write(requestInfoBytes)
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
