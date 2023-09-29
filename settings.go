package settings

import (
	"fmt"
	"log"
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
