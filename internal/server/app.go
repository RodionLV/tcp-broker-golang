package server

import (
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

	queueMap := make(map[string]chan []byte)

	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("client is connected", client)
		go handleConnection(client, queueMap)
	}
}
