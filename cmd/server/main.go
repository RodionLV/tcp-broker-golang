package main

import (
	"tcp-broker-golang/internal/server"
)

// todo: add config
// todo: fix send in queue only two message, other will delete

func main() {
	server.Run()
}
