package client

import (
	"fmt"
	"net"
	"tcp-broker-golang/pkg/tcp/transfer"
)

func handleSending(client net.Conn) {
	var message string

	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&message)
		if err != nil {
			continue
		}

		_, err = transfer.WriteMessage(client, message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func handleReceiving(client net.Conn) {
	for {
		data, err := transfer.ReadMessage(client)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("receive:", string(data))
	}
}
