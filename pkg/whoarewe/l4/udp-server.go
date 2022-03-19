package l4

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
	"sync"
)

func StartUDPEchoServer(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Starting UDP server at:", viper.GetString("UDP_PORT"))
	udpListener, err := net.ListenPacket("udp", "0.0.0.0:"+viper.GetString("UDP_PORT"))
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	defer udpListener.Close()
	for {
		buffer := make([]byte, 65508) // max range + 1, I don't know if they append '\n too'
		n, clientAddress, err := udpListener.ReadFrom(buffer)
		if err != nil {
			log.Printf("Unable to read from new connection %v : %v\n", clientAddress, err)
			continue
		}
		requestData := string(buffer[:n])
		log.Printf("Received UDP request from client %v : %v\n", clientAddress, requestData)
		if n > 0 {
			go handleNewUDPConnection(&requestData, &clientAddress, &udpListener)
		}
	}
}

func handleNewUDPConnection(requestData *string, clientAddress *net.Addr, udpListener *net.PacketConn) {
	response := fmt.Sprintf(
		"\"%v\" serving UDP at %v:%v from \"%v\" : %v\n",
		os.Getenv("POD_NAME"), "0.0.0.0", viper.GetString("UDP_PORT"), os.Getenv("POD_NAMESPACE"), *requestData)

	_, err := (*udpListener).WriteTo([]byte(response), *clientAddress)
	if err != nil {
		log.Println("Error writing to connection: ", err.Error())
		return
	}
	//_, err = (*udpListener).WriteTo([]byte("FYI, that address can be used back anytime!!\n"), *clientAddress)
	//if err != nil {
	//	log.Println("Error writing to connection: ", err.Error())
	//	return
	//}
}
