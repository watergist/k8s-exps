package regsiter

import (
	"github.com/watergist/k8s-manifests/pkg/whoarewe/listener"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/whoami"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/whoareu"
	"net/http"
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
