package app

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"hawx.me/code/mpd-scrobbler/config"
)

// ParseFile will unmarshall the JSON file to check if it fits with the
// Settings struct in config package
func ParseFile(file string) config.Settings {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)

	var settings config.Settings
	err := json.Unmarshal([]bye(s), &settings)
	if err != nil {
		log.Fatal(err)
	}
	return settings
}
