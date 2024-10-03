package client

import (
	"fmt"
	"net"
)

func Run() {
	fmt.Println("client is running")

	server, err := net.Dial("tcp", "127.0.0.1:4001")

	if err != nil {
		fmt.Println(err)
		return
	}

	var message string
	for {
		fmt.Print("write: ")
		_, err := fmt.Scan(&message)

		if err != nil {
			fmt.Println(err)
			continue
		}

		server.Write([]byte(message))
	}

	defer server.Close()
}
