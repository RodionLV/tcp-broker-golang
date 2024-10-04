package client

import (
	"fmt"
	"net"
	"tcp-broker-golang/pkg/tcp/transfer"
)

func Run() {
	fmt.Println("client is running")

	client, err := net.Dial("tcp", "127.0.0.1:4001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	typeClient := selectingTypeClient()

	_, err = client.Write([]byte(typeClient + "\n"))

	if err != nil {
		fmt.Println(err)
		return
	}

	switch typeClient {
	case "send":
		handleSending(client)
	case "receive":
		handleReceiving(client)
	}
}

func selectingTypeClient() string {
	var typeClient string

	fmt.Println("write type (send/receive):")
	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&typeClient)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if typeClient == "send" || typeClient == "receive" {
			return typeClient
		}
	}
}

func handleSending(client net.Conn) {
	fmt.Println("sendler:")

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
	fmt.Println("receiver:")
}
