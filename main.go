package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":6667")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Client connected")
		go handleConnection(conn)
	}

}

func handleConnection(c net.Conn) {
	io.Copy(c, c)
	c.Close()
}
