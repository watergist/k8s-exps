package whoami

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func podName(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("I am %v\n", os.Getenv("POD_NAME"))))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func podInfo(w http.ResponseWriter, _ *http.Request) {
	file, err := os.ReadFile("/config/pod-info/labels")
	if err != nil {
		writeServerError(w, err, "reading labels", http.StatusNotImplemented)
		return
	}
	_, err = w.Write([]byte(file))
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
