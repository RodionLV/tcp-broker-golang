package server

import (
	"fmt"
	"net"
	"tcp-broker-golang/pkg/tcp/transfer"
)

func handleConnection(client net.Conn, queueCh chan []byte) {
	defer client.Close()

	typeClient, err := initTypeClient(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch typeClient {
	case "sendler":
		handleSendler(client, queueCh)
	case "receiver":
		handleReceiver(client, queueCh)
	}
}

func handleReceiver(client net.Conn, queueCh chan []byte) {
	for {
		data := Get(queueCh)

		transfer.WriteMessage(client, string(data))
	}
}

func handleSendler(client net.Conn, queueCh chan []byte) {
	for {
		data, err := transfer.ReadMessage(client)
		if err != nil {
			continue
		}

		go Put(queueCh, data)

		fmt.Println("put message:", string(data))
	}
}

func Put(queueCh chan []byte, data []byte) {
	queueCh <- data
}

func Get(queueCh chan []byte) []byte {
	return <-queueCh
}
