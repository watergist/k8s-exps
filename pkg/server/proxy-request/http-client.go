package proxy_request

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type upstreamRequest struct {
	errorPlace  string
	upstreamErr error
	body        []byte
	resp        *http.Response
	reqUrl      string
}

func request(reqUrl url.URL, requestMethod string) (upstreamRequestData *upstreamRequest) {
	upstreamRequestData = &upstreamRequest{reqUrl: reqUrl.String()}
	client := http.Client{
		Timeout: time.Duration(viper.GetInt("UPSTREAM_TIMEOUT")),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	switch requestMethod {
	case "GET":
		upstreamRequestData.resp, upstreamRequestData.upstreamErr = client.Get(reqUrl.String())
	}

	if upstreamRequestData.upstreamErr != nil {
		upstreamRequestData.errorPlace = "error making connection to upstream"
		log.Printf(upstreamRequestData.errorPlace, upstreamRequestData.upstreamErr)
		return
	}
	upstreamRequestData.body, upstreamRequestData.upstreamErr = ioutil.ReadAll(upstreamRequestData.resp.Body)
	if upstreamRequestData.upstreamErr != nil {
		upstreamRequestData.errorPlace = "error reading data from upstream: "
		log.Printf(upstreamRequestData.errorPlace, upstreamRequestData.upstreamErr)
		return
	}
	return
}

func makeReq(w *http.ResponseWriter, upstreamRequestData *upstreamRequest) {
	(*w).Header().Add("X-base-http-upstream-url", upstreamRequestData.reqUrl)
	if upstreamRequestData.upstreamErr != nil {
		(*w).Header().Add("X-base-http-upstream-conn-err", upstreamRequestData.upstreamErr.Error())
		(*w).Header().Add("X-base-http-upstream-conn-error-place", upstreamRequestData.errorPlace)
	}
	(*w).WriteHeader(upstreamRequestData.resp.StatusCode)
	_, err := (*w).Write(upstreamRequestData.body)
	if err != nil {
		log.Print("error writing data to downstream: ", err)
		return
	}
}
