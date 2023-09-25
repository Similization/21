package main

import (
	"log"
	"net"
)

func main() {
	message := "Hi, I am server"

	listener, err := net.Listen("tcp", "127.0.0.1:4545")
	if err != nil {
		log.Fatal("Couldn't create listener: ", err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("Couldn't accept connection: ", err)
		}
		defer connection.Close()

		connection.Write([]byte(message))
	}
}
