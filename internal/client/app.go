package client

import (
	"fmt"
	"net"
	"strings"
)

const DELIM = ":"

func Run() {
	fmt.Println("client is running")

	client, err := net.Dial("tcp", "127.0.0.1:4001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	typeClient, topic := selectingTypeClient()
	_, err = client.Write([]byte(typeClient + DELIM + topic + "\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	switch typeClient {
	case "send":
		fmt.Println("sending...")

		handleSending(client)
	case "receive":
		fmt.Println("receiving...")

		handleReceiving(client)
	}
}

func selectingTypeClient() (string, string) {
	var initData string

	fmt.Println("write type (send/receive) and topic (send:topic):")
	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&initData)
		if err != nil {
			fmt.Println(err)
			continue
		}

		values := strings.Split(initData, DELIM)
		if len(values) != 2 {
			continue
		}

		if values[0] == "send" || values[0] == "receive" {
			return values[0], values[1]
		}
	}
}
