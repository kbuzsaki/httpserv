package main

import (
	"github.com/kbuzsaki/httpserv/server"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server.ServeEcho(listener)
}
