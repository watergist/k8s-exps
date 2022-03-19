package main

import (
	"github.com/spf13/viper"
	"github.com/watergist/k8s-manifests/pkg/whoarewe/l4"
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
	go l4.StartTCPEchoServer(&wg)
	go l4.StartUDPEchoServer(&wg)
	wg.Wait()
	log.Println("Application Terminated")
}
