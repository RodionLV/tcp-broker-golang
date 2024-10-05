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

		transfer.WriteMessage(client, message)
	}
}

func handleReceiving(client net.Conn) {
	for {
		data, err := transfer.ReadMessage(client)
		if err != nil {
			continue
		}

		fmt.Println("receive:", string(data))
	}
}
