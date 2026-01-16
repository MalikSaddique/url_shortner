package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	URLShortner struct {
		POSTGRES_HOST     string
		POSTGRES_PORT     string
		POSTGRES_USER     string
		POSTGRES_PASSWORD string
		POSTGRES_DB       string
	}
}

var Cfg Config

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	cfg.URLShortner.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	cfg.URLShortner.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	cfg.URLShortner.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	cfg.URLShortner.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	cfg.URLShortner.POSTGRES_DB = os.Getenv("POSTGRES_DB")

	Cfg = cfg
	return &cfg, nil
}
