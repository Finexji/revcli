// reverse-server.go
package main

import (
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	// Wait for incoming connection from client-side to pipe
	target, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		log.Println("Target dial error:", err)
		return
	}
	defer target.Close()

	go io.Copy(conn, target)
	io.Copy(target, conn)
}

func main() {
	listener, err := net.Listen("tcp", ":6324")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	log.Println("Listening on :6324")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go handle(conn)
	}
}
