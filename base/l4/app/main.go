package main

import (
	"github.com/spf13/viper"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/l4"
	"log"
	"time"
)

func main() {
	log.Println("Started Application")
	viper.AutomaticEnv()
	viper.SetDefault("TCP_PORT", "31400")
	l4.StartTcpEchoServer()
	time.Sleep(10000 * time.Second)
	log.Println("Exited Application")
}
