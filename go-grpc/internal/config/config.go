package config

import (
	"log"
	"os"
)

type mysql struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	Port  string
	Mysql mysql
}

func GetConfig() Config {
	log.Println("Initializing config...")

	// switch service mode if needed
	return Config{
		Port: os.Getenv("SERVICE_PORT"),
		Mysql: mysql{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Database: os.Getenv("MYSQL_DATABASE"),
		},
	}
}
