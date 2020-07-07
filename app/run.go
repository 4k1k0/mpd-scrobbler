package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"hawx.me/code/mpd-scrobbler/client"
	"hawx.me/code/mpd-scrobbler/config"
	"hawx.me/code/mpd-scrobbler/scrobble"
)

const (
	// only submit tracks longer then minTrackLen
	minTrackLen = 30

	// only submit if played for submitTime second or submitPercentage of length
	submitTime       = 240
	submitPercentage = 50

	// polling interval
	sleepTime = 5 * time.Second
)

func catchInterrupt() {
	// To allow the user to send Ctrl-C to shut down the program
	c := make(chan os.Signal, 1)
	// This will shut down the program when channel c gets at least 1 signal
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	log.Printf("caught %s: shutting down", s)
}

// StartApp runs the program this way the main package is clean
func StartApp() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file")
		panic(err)
	}
	settings := config.New()

	// Connect to MPD
	c, err := client.Dial("tcp", fmt.Sprintf(":%d", settings.Mpd.Port))
	if err != nil {
		log.Fatal(err)
	}
	// Close the conection at the end of the main goroutine
	defer c.Close()

	// It creates a database connection?
	// and defer the database close
	db, err := scrobble.Open(settings.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TODO: Handle this hardcoded strings
	api, err := scrobble.New(db, "lastfm", settings.User.Key, settings.User.Secret, settings.User.Username, settings.User.Password, "")
	if err != nil {
		log.Fatal(err)
	}

	// This channels stores the song to submit and the now playing values
	toSubmit := make(chan client.Song)
	nowPlaying := make(chan client.Song)

	go c.Watch(sleepTime, toSubmit, nowPlaying)

	go func() {
		for {
			select {
			case s := <-nowPlaying:
				err := api.NowPlaying(s.Artist, s.Album, s.AlbumArtist, s.Title)
				if err != nil {
					log.Printf("[%s] err(NowPlaying): %s\n", api.Name(), err)
				}
			case s := <-toSubmit:
				err := api.Scrobble(s.Artist, s.Album, s.AlbumArtist, s.Title, s.Start)
				if err != nil {
					log.Printf("[%s] err(Scrobble): %s\n", api.Name(), err)
				}
			}
		}
	}()

	catchInterrupt()
}
