package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

type serverStruct struct {
	host string
	port string
}

func (s *serverStruct) toString() string {
	return fmt.Sprintf("%s:%s", s.host, s.port)
}

func readFromEnv() *serverStruct {
	return &serverStruct{
		host: getEnv("HOST", ""),
		port: getEnv("PORT", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Print("No variable named " + key + " was found.")
	return defaultVal
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	address := readFromEnv().toString()
	connection, err := net.Dial("tcp", address)
	if err != nil {
		log.Panic("Couldn't connect to server: ", err)
	}
	defer connection.Close()

	for {
		for {
			input := make([]byte, 1024)
			n, err := connection.Read(input)

			if err != nil {
				log.Print("Couldn't read from server: ", err)
			}
			fmt.Printf("Server: %s\n", string(input[:n]))

			if n == 0 {
				break
			}
		}
		var answer string
		fmt.Scan(&answer)
		connection.Write([]byte(answer))
	}
}
