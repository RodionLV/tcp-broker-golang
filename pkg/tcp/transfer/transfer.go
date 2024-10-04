package transfer

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

func ReadMessage(client net.Conn) ([]byte, error) {
	var clientReader = bufio.NewReader(client)

	msg, err := clientReader.ReadString(' ')
	if err != nil {
		return nil, err
	}

	size, err := strconv.Atoi(strings.Trim(msg, " "))
	if err != nil {
		return nil, err
	}

	var data []byte = make([]byte, size)
	clientReader.Read(data)

	return data, nil
}

func WriteMessage(client net.Conn, message string) (int, error) {
	return client.Write([]byte(strconv.Itoa(len(message)) + " " + message))
}
