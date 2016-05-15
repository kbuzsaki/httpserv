package server

import (
	"github.com/kbuzsaki/httpserv/http"
	"log"
	"net"
)

type Handler interface {
	Handle(request http.Request) http.Response
}

func ServeHandler(listener net.Listener, handler Handler) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("err accepting:", err)
		}

		go HandleRequest(conn, handler)
	}
}

func HandleRequest(conn net.Conn, handler Handler) {
	defer conn.Close()

	// read and parse header
	request, err := http.ReadRequest(conn)
	if err != nil {
		log.Println(err)
		return
	}

	// process request with handler
	response := handler.Handle(request)

	// write out response
	conn.Write([]byte(response.Protocol.String() + " "))
	conn.Write([]byte(response.Status.String() + "\n"))
	conn.Write([]byte("\n"))
	conn.Write([]byte(response.Body + "\n"))
}

// sample hello world server and handler
func ServeHelloWorld(listener net.Listener) {
	ServeHandler(listener, HelloWorldHandler{})
}

// sample echoing server and handler
func ServeEcho(listener net.Listener) {
	ServeHandler(listener, EchoHandler{})
}
