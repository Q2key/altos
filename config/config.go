package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort       string
	ServerHost       string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	return &Config{
		ServerPort:       os.Getenv("SERVER_PORT"),
		ServerHost:       os.Getenv("SERVER_HOST"),
		DatabasePort:     os.Getenv("DB_PORT"),
		DatabaseUser:     os.Getenv("DB_USER"),
		DatabasePassword: os.Getenv("DB_PASSWORD"),
		DatabaseName:     os.Getenv("DB_NAME"),
	}
}
