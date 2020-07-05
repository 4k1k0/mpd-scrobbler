package app

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"hawx.me/code/mpd-scrobbler/client"
	"hawx.me/code/mpd-scrobbler/config"
)

// StartApp runs the program this way the main package is clean
func StartApp() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file")
		panic(err)
	}
	settings := config.New()

	// Connect to MPD
	c, err := client.Dial("tcp", fmt.Sprintf(":%d", settings.Port))
	if err != nil {
		log.Fatal(err)
	}
	// Close the conection at the end of the main goroutine
	defer c.Close()
}
