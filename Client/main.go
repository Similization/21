package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		log.Panic("Couldn't connect to server: ", err)
	}
	defer connection.Close()

	input := make([]byte, 1024)
	n, err := connection.Read(input)
	if err != nil {
		log.Print("Couldn't read from server: ", err)
	}

	fmt.Printf("%s\n%d bytes have been read", string(input[:n]), n)
}
