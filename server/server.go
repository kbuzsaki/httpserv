package server

import (
	"github.com/kbuzsaki/http/http"
	"log"
	"net"
)

func ServeHelloWorld(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("err accepting:", err)
		}

		go HandleHelloWorld(conn)
	}
}

func HandleHelloWorld(conn net.Conn) {
	// read header
	request, err := http.ReadRequest(conn)
	if err != nil {
		log.Println(err)
		return
	}

	// write sample header
	conn.Write([]byte(http.HttpOneDotOne.String() + " "))
	conn.Write([]byte(http.StatusOk.String() + "\n"))
	conn.Write([]byte("\n"))

	// write sample body
	conn.Write([]byte("<h1>Hello World</h1>\n"))
	conn.Write([]byte("<p>You requested path: " + request.Path + "</p>\n"))

	conn.Write([]byte("<table><thead><th>Key</th><th>Value</th></thead><tbody>"))
	for _, param := range request.Query {
		conn.Write([]byte("<tr><td>" + param.Key + "</td><td>" + param.Val + "</td></tr>\n"))
	}
	conn.Write([]byte("</tbody></table>\n"))

	conn.Close()
}
