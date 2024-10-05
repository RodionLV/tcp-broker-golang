package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"tcp-broker-golang/pkg/tcp/transfer"
)

const DELIM = ":"

func handleConnection(client net.Conn, queueMap map[string]chan []byte) {
	defer client.Close()

	typeClient, topic, err := initTypeClient(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, isExists := queueMap[topic]
	if !isExists {
		fmt.Println("creating queue named:", topic)
		queueMap[topic] = make(chan []byte)
	}

	fmt.Println("setup\n client type:", typeClient, "\n queue:", topic)

	clientSetup := ClientSetup{
		Client:     client,
		QueueCh:    queueMap[topic],
		ClientType: typeClient,
	}

	switch typeClient {
	case "sendler":
		handleSendler(clientSetup)
	case "receiver":
		handleReceiver(clientSetup)
	}
}

func handleReceiver(clientSetup ClientSetup) {
	for {
		data := clientSetup.Get()
		_, err := transfer.WriteMessage(clientSetup.Client, string(data))

		if err != nil {
			fmt.Println(err)
			go clientSetup.Put(data)
			break
		}
	}
}

func handleSendler(clientSetup ClientSetup) {
	for {
		data, err := transfer.ReadMessage(clientSetup.Client)
		if err != nil {
			fmt.Println(err)
			return
		}

		go clientSetup.Put(data)
		fmt.Println("put message:", string(data))
	}
}

func initTypeClient(client net.Conn) (string, string, error) {
	for {
		msg, err := bufio.NewReader(client).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return "", "", err
		}

		if len(msg) <= 3 {
			continue
		}

		values := strings.Split(msg[:len(msg)-1], DELIM)
		if len(values) != 2 {
			continue
		}

		switch values[0] {
		case "send":
			return "sendler", values[1], nil
		case "receive":
			return "receiver", values[1], nil
		default:
			fmt.Println("message '", strings.Join(values, DELIM), "' is wrong")
		}
	}
}
