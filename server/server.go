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
	header, err := http.ReadHeader(conn)
	if err != nil {
		log.Println(err)
		return
	}

	// write sample header
	conn.Write([]byte(http.HttpOneDotOne.String() + " "))
	conn.Write([]byte(http.StatusOk.String() + "\n"))
	conn.Write([]byte("\n"))

	// write sample body
	conn.Write([]byte("Hello World\n"))
	conn.Write([]byte(header.String() + "\n"))

	conn.Close()
}
