package main

import (
	"github.com/spf13/viper"
	multiserver "github.com/watergist/k8s-manifests/pkg/whoarewe/multi-server"
	"log"
)

func main() {
	log.Println("Application Started")
	viper.AutomaticEnv()
	//viper.SetDefault("HTTP_PORTS", 1080)
	//viper.SetDefault("HTTPS_PORTS", 1443)
	//viper.SetDefault("TLS_KEYPAIR_PATH", "base/http/app/testdata")
	s := multiserver.GetServer()
	if viper.GetString("TLS_KEYPAIR_PATH") == "" {
		registerLocalCertificate(s)
	}
	s.RunServers()
	s.WG.Wait()
	log.Println("Application Terminated")
}
