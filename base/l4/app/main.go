package main

import (
	"github.com/spf13/viper"
	l42 "github.com/watergist/k8s-manifests/pkg/server/whoarewe/l4"
	"log"
	"sync"
)

func main() {
	log.Println("Application Started")
	viper.AutomaticEnv()
	//viper.SetDefault("TCP_PORT", "31400")
	//viper.SetDefault("UDP_PORT", "31400")
	wg := sync.WaitGroup{}
	wg.Add(2)
	go l42.StartTCPEchoServer(&wg)
	go l42.StartUDPEchoServer(&wg)
	wg.Wait()
	log.Println("Application Terminated")
}
