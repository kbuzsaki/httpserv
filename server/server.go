package server

import (
	"github.com/kbuzsaki/http/http"
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

type HelloWorldHandler struct {
}

func (h HelloWorldHandler) Handle(request http.Request) http.Response {
	return http.Response{http.HttpOneDotOne, http.StatusOk, "Hello World"}
}

// sample echoing server and handler
func ServeEcho(listener net.Listener) {
	ServeHandler(listener, EchoHandler{})
}

type EchoHandler struct {
}

func (h EchoHandler) Handle(request http.Request) http.Response {
	var response http.Response

	response.Protocol = http.HttpOneDotOne
	response.Status = http.StatusOk

	body := "<h1>Echo Response</h1>\n"
	body += "<p>You requested path: " + request.Path + "</p>\n"

	body += "<table><thead><th>Key</th><th>Value</th></thead><tbody>"
	for _, param := range request.Query {
		body += "<tr><td>" + param.Key + "</td><td>" + param.Val + "</td></tr>\n"
	}
	body += "</tbody></table>\n"

	response.Body = body

	return response
}
