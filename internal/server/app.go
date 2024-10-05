package server

import (
	"bufio"
	"fmt"
	"net"
)

func Run() {
	fmt.Println("Server is running")

	listener, err := net.Listen("tcp", ":4001")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	var queueCh = make(chan []byte)

	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("client is connected", client)

		go handleConnection(client, queueCh)
	}
}

func initTypeClient(client net.Conn) (string, error) {
	for {
		msg, err := bufio.NewReader(client).ReadString('\n')
		if err != nil {
			return "", err
		}

		if len(msg) <= 1 {
			continue
		}

		switch msg[:len(msg)-1] {
		case "send":
			return "sendler", nil
		case "receive":
			return "receiver", nil
		default:
			fmt.Println("message '", msg[:len(msg)-1], "' is wrong")
		}
	}
}
