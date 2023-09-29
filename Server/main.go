package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var deck = []string{
	"6 of Hearts", "6 of Diamonds", "6 of Clubs", "6 of Spades",
	"7 of Hearts", "7 of Diamonds", "7 of Clubs", "7 of Spades",
	"8 of Hearts", "8 of Diamonds", "8 of Clubs", "8 of Spades",
	"9 of Hearts", "9 of Diamonds", "9 of Clubs", "9 of Spades",
	"10 of Hearts", "10 of Diamonds", "10 of Clubs", "10 of Spades",
	"Jack of Hearts", "Jack of Diamonds", "Jack of Clubs", "Jack of Spades",
	"Queen of Hearts", "Queen of Diamonds", "Queen of Clubs", "Queen of Spades",
	"King of Hearts", "King of Diamonds", "King of Clubs", "King of Spades",
	"Ace of Hearts", "Ace of Diamonds", "Ace of Clubs", "Ace of Spades",
}

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

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Couldn't create listener: ", err)
	}
	defer listener.Close()
	fmt.Println("Server is running ...")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("Couldn't accept connection: ", err)
		}
		defer connection.Close()

		connection.Write([]byte("List of commands:\n1. Start\n2. Cards\n3. Pick\n4. Pass\n5. End\n"))

		var newDeck, userCards []string
		input := make([]byte, 1024)
		for {
			n, err := connection.Read(input)
			if err != nil {
				log.Print(err)
			}

			res := strings.ToLower(string(input[:n]))
			switch res {
			case "start":
				newDeck = shuffleDeck(deck)
				userCards = make([]string, 2)
				connection.Write([]byte("Game was created"))
				fallthrough
			case "pick":
				card := getRandomElementOfDeck(newDeck)
				userCards = append(userCards, card)
				connection.Write([]byte("You take: " + card))
			case "cards":
				connection.Write([]byte("Cards: " + strings.Join(userCards, "\n")))
			case "pass":
				connection.Write([]byte("Appears in the future"))
			case "end":
				connection.Write([]byte("Appears in the future"))
			}
		}
	}
}

func shuffleDeck(deck []string) []string {
	// Make a copy of the deck
	shuffledDeck := make([]string, len(deck))
	copy(shuffledDeck, deck)

	// Seed the random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Fisher-Yates shuffle algorithm
	for i := len(shuffledDeck) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	}

	return shuffledDeck
}

func getRandomElementOfDeck(deck []string) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	i := r.Intn(len(deck))
	el := deck[i]
	deck = append(deck[:i], deck[i+1:]...)
	return el
}
