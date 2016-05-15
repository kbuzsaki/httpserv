package server

import (
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
	var buf []byte
	var n int
	var err error

	log.Println("request contents:")
	n, err = conn.Read(buf)
	for err == nil {
		log.Println(string(buf[:n]))
		n, err = conn.Read(buf)
	}

	resp := []string{"HTTP/1.1 200 OK", "", "Hello World"}

	for _, line := range resp {
		conn.Write([]byte(line + "\n"))
	}

	log.Println("response complete")

	conn.Close()
}
