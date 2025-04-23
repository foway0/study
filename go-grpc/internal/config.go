package internal

import (
	"log"
	"os"
)

type Config struct {
	Port string
}

var config *Config

func GetConfig() *Config {
	if config != nil {
		log.Println("Config already initialized")
		return config
	}

	log.Println("Initializing config...")

	// switch service mode if needed
	config = &Config{
		Port: os.Getenv("SERVICE_PORT"),
	}

	return config
}
