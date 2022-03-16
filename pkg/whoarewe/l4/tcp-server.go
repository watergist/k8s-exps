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
)

func StartTcpEchoServer() {
	log.Println("Starting tcp server at:", viper.GetString("TCP_PORT"))
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
		go handleNewConnection(conn)
	}
}

// Concurrently handle one tcp connection
func handleNewConnection(conn net.Conn) {
	defer conn.Close()
	connectionReader := bufio.NewReader(conn)
	var requestData, requestDataLine string
	var err error
	for {
		requestDataLine, err = connectionReader.ReadString('\n')
		if len(requestDataLine) > 0 {
			response := fmt.Sprintf(
				"\"%v\" serving at %v:%v from \"%v\" : %v\n",
				os.Getenv("POD_NAME"), "0.0.0.0", viper.GetString("TCP_PORT"), os.Getenv("POD_NAMESPACE"), requestDataLine)
			_, err = conn.Write([]byte(response))
			if err != nil {
				log.Println("Error writing request: ", err.Error())
				return
			}
		}
		if err == nil {
			continue
		} else if errors.Is(err, io.EOF) {
			log.Printf("Received request from client %v : %v\n", conn.RemoteAddr(), requestData)
			break
		} else if err != nil {
			log.Println("Error reading request: ", err.Error())
			break
		}
	}
}
