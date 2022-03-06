package wide

import (
	"bufio"
	"crypto/tls"
	"net/http"
	"net/url"
)

type RequestProperties struct {
	ClientAddress        string
	Request              map[string]string
	URL                  *url.URL
	QueryParams          *url.Values
	Headers              *http.Header
	TLS                  *tls.ConnectionState
	BasicResponseHeaders *http.Header
	Body                 []byte
	Error                []error
}

func GenRequestInfo(r *http.Request) RequestProperties {
	err := r.ParseForm()
	requestReader := bufio.NewReader(r.Body)
	requestBytes, err2 := requestReader.Peek(requestReader.Size())
	requestProperties := RequestProperties{
		ClientAddress: r.RemoteAddr,
		Request: map[string]string{
			"host":   r.Host,
			"uri":    r.RequestURI,
			"method": r.Method,
		},
		URL:         r.URL,
		QueryParams: &r.Form,
		Headers:     &r.Header,
		Body:        requestBytes,
		TLS:         r.TLS,
		Error:       nil,
	}
	if err != nil {
		requestProperties.Error = append(requestProperties.Error, err)
	}
	if err2 != nil {
		requestProperties.Error = append(requestProperties.Error, err2)
	}
	return requestProperties
}
