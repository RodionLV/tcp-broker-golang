package server

import "net"

type ClientSetup struct {
	ClientType string
	QueueCh    chan []byte
	Client     net.Conn
}

func (c *ClientSetup) Put(data []byte) {
	c.QueueCh <- data
}

func (c *ClientSetup) Get() []byte {
	return <-c.QueueCh
}
