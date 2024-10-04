package server

import (
	"bufio"
	"fmt"
	"net"
	"tcp-broker-golang/pkg/tcp/transfer"
)

func Run() {
	fmt.Println("Server is running")

	listener, err := net.Listen("tcp", ":4001")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("client is connected", client)

		go handleConnection(client)
	}
}

func handleConnection(client net.Conn) {
	defer client.Close()

	var typeClient string

	for {
		if len(typeClient) == 0 {
			msg, err := bufio.NewReader(client).ReadString('\n')
			if err != nil {
				return
			}

			if len(msg) <= 1 {
				continue
			}

			switch msg[:len(msg)-1] {
			case "send":
				typeClient = "sendler"
			case "receive":
				typeClient = "receiver"
			default:
				fmt.Println("message '", msg[:len(msg)-1], "' is wrong")
			}
		} else {
			data, err := transfer.ReadMessage(client)
			if err != nil {
				continue
			}

			fmt.Println("message:", string(data), "from:", typeClient)
		}
	}
}

func Put() {

}

func Get() {

}
