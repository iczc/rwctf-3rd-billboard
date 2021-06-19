package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	CheckMode string
	Flag      string
	LCD       string
	Port      string
}

func NewConfig() *Config {
	return &Config{
		CheckMode: os.Getenv("CHECK_MODE"),
		Flag:      os.Getenv("FLAG"),
		LCD:       os.Getenv("LCD"),
		Port:      os.Getenv("PORT"),
	}
}
