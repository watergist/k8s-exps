package proxy_request

import (
	"github.com/watergist/k8s-manifests/pkg/server"
	"github.com/watergist/k8s-manifests/pkg/server/listener"
	"net/http"
	"net/url"
)

type Server struct {
	*listener.Server
}

func (s *Server) Proxy(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		server.WriteServerError(w, err, "parsing query params", 401)
	}
	formData := r.Form
	req := url.URL{
		Scheme: formData.Get("scheme"),
		Host:   formData.Get("host"),
		Path:   formData.Get("uri"),
	}
	if req.Host == "" {
		req.Host = r.Host
	}
	if req.Path == "" {
		req.Path = "/"
	}
	if req.Scheme == "" {
		req.Scheme = "http"
	}

	reqMethod := formData.Get("req-method")
	if reqMethod == "" {
		reqMethod = "GET"
	}

	upstreamRequestData := request(req, reqMethod)
	makeReq(&w, upstreamRequestData)
}
