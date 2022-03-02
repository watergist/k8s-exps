package whoami

import "net/http"

func RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/whoami", podName)
	mux.HandleFunc("/whoami/wide", podInfo)
}
