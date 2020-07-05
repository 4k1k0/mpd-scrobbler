package config

import (
	"os"
	"strconv"
)

// config file
// database file
// mpd port
// mpd server ?

// Settings contains the information to succesfuly config a
// Last.fm user for MPD
type Settings struct {
	Database string
	Port     int
	Server   string
	User     User
}

// User contains username and password for
// your Last.fm account
type User struct {
	Username string
	Password string
	Key      string
	Secret   string
}

// New returns a pointer to a settings struct
func New() *Settings {
	return &Settings{
		Database: getEnvAsString("DATABASE", "./scrobble.db"),
		Port:     getEnvAsInt("PORT", 6600),
		Server:   getEnvAsString("SERVER", "localhost"),
		User: User{
			Username: getEnvAsString("USERNAME", ""),
			Password: getEnvAsString("PASSWORD", ""),
			Key:      getEnvAsString("KEY", ""),
			Secret:   getEnvAsString("SECRET", ""),
		},
	}
}

// getEnv returns the value of a environment variable or a default value as a string
func getEnvAsString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnv returns the value of a environment variable or a default value as a int
func getEnvAsInt(key string, defaultValue int) int {
	valueString := getEnvAsString(key, "")
	if value, err := strconv.Atoi(valueString); err == nil {
		return value
	}
	return defaultValue
}
