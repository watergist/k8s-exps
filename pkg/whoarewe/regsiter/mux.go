package regsiter

import (
	"github.com/watergist/k8s-manifests/pkg/whoarewe/listener"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/whoami"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/whoareu"
	"log"
	"net/http"
	"os"
	"time"
)

func RegisterEndpoints(mux *http.ServeMux, listenerProperties *listener.Listener) {
	whoareuEndpoints := whoareu.Server{ListenerProperties: listenerProperties}
	whoamiEndpoints := whoami.Server{ListenerProperties: listenerProperties}
	mux.HandleFunc("/whoareu", whoareuEndpoints.PodName)
	mux.HandleFunc("/whoareu/wide", whoareuEndpoints.PodInfo)
	mux.HandleFunc("/whereami", whoamiEndpoints.RequestHost)
	mux.HandleFunc("/whoami", whoamiEndpoints.RequestIP)
	mux.HandleFunc("/whoami/wide", whoamiEndpoints.RequestInfo)
}

func EnableDualLogging(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("request for %v\n", r.URL.Path)
			start := time.Now()
			w.Header().Add("a-forwArDed-by", "whoarewe")
			w.Header().Add("a-forwArDed-by", os.Getenv("POD_NAME"))
			mux.ServeHTTP(w, r)
			log.Printf("served %v in %v milliseconds", r.URL.Path, time.Now().Sub(start).Milliseconds())
		},
	)
}
