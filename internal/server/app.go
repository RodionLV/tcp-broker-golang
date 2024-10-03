package server

import (
	"fmt"
	"io"
	"net"
	"os"
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

		go func(client net.Conn) {
			defer client.Close()

			for {
				io.Copy(os.Stdout, client)
				fmt.Println("get message")
			}
		}(client)
	}
}
