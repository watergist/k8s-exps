package whoarewe

import (
	"net/http"
	"watergist/k8s-manifests/pkg/whoarewe/whoami"
	"watergist/k8s-manifests/pkg/whoarewe/whoareu"
)

func RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/whoareu", whoareu.PodName)
	mux.HandleFunc("/whoareu/wide", whoareu.PodInfo)
	mux.HandleFunc("/whereami", whoami.RequestHost)
	mux.HandleFunc("/whoami", whoami.RequestIP)
	mux.HandleFunc("/whoami/wide", whoami.RequestInfo)
}
