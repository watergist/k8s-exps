package l4

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func StartTCPEchoServer(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Starting TCP server at:", viper.GetString("TCP_PORT"))
	tcpListener, err := net.Listen("tcp", "0.0.0.0:"+viper.GetString("TCP_PORT"))
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	defer tcpListener.Close()
	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		log.Printf("Received TCP request from client %v\n", conn.RemoteAddr())
		go handleNewTCPConnection(conn)
	}

}

func handleNewTCPConnection(conn net.Conn) {
	defer conn.Close()
	connectionReader := bufio.NewReader(conn)
	var requestDataLine string
	var err error
	for {
		requestDataLine, err = connectionReader.ReadString('\n')
		log.Printf("Received TCP data from client %v : %v\n", conn.RemoteAddr(), requestDataLine)
		if len(requestDataLine) > 0 {
			response := fmt.Sprintf(
				"\"%v\" serving TCP at %v:%v from \"%v\" : %v\n",
				os.Getenv("POD_NAME"), "0.0.0.0", viper.GetString("TCP_PORT"), os.Getenv("POD_NAMESPACE"), requestDataLine)
			_, err = conn.Write([]byte(response))
			if err != nil {
				log.Println("Error writing to connection: ", err.Error())
				return
			}
		}
		if errors.Is(err, io.EOF) {
			log.Printf("Request from client completed %v : %v\n", conn.RemoteAddr(), requestDataLine)
			break
		} else if err != nil {
			log.Println("Error reading request: ", err.Error())
			break
		}
	}
}
