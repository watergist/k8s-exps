package multiserver

import (
	"github.com/spf13/viper"
	"github.com/watergist/k8s-manifests/pkg/server/listener"
	"github.com/watergist/k8s-manifests/pkg/server/proxy-request"
	"github.com/watergist/k8s-manifests/pkg/server/whoarewe/whoami"
	"github.com/watergist/k8s-manifests/pkg/server/whoarewe/whoareu"
	"log"
	"net/http"
	"os"
	"time"
)

func RegisterEndpoints(mux *http.ServeMux, listenerProperties *listener.Listener) {
	whoareuEndpoints := whoareu.Server{&listener.Server{ListenerProperties: listenerProperties}}
	whoamiEndpoints := whoami.Server{&listener.Server{ListenerProperties: listenerProperties}}
	proxyRequestEndpoints := proxy_request.Server{&listener.Server{ListenerProperties: listenerProperties}}

	mux.HandleFunc("/whoareu", whoareuEndpoints.PodName)
	mux.HandleFunc("/whoareu/wide", whoareuEndpoints.PodInfo)
	mux.HandleFunc("/whoami", whoamiEndpoints.RequestIP)
	mux.HandleFunc("/whoami/wide", whoamiEndpoints.RequestInfo)
	mux.HandleFunc("/", whoamiEndpoints.RequestHost)
	mux.HandleFunc("/proxy-request", proxyRequestEndpoints.Proxy)
}

func EnableDualLogging(mux *http.ServeMux, address string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("request at %v for %v\n", address, r.URL.Path)
			start := time.Now()
			time.Sleep(time.Millisecond * time.Duration(viper.GetInt("SERVER_FAULT_DELAY")))
			w.Header().Add("a-forwArDed-by", os.Getenv("POD_NAME"))
			w.Header().Add("a-FORWArDed-by", os.Getenv("POD_NAMESPACE"))
			mux.ServeHTTP(w, r)
			log.Printf("served at %v for %v in %v milliseconds", address, r.URL.Path, time.Now().Sub(start).Milliseconds())
		},
	)
}
