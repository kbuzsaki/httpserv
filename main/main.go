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

	routes := []server.Route{
		{"/", server.HelloWorldHandler{}},
		{"/hello", server.HelloWorldHandler{}},
	}
	handler := server.RoutingHandler{
		routes,
		server.EchoHandler{},
	}
	server.ServeHandler(listener, &handler)
}
