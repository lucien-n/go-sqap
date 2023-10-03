package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string

	APIHost string
	APIPort string
}

func LoadConfig(dotenv_path string) *Config {
	err := godotenv.Load(dotenv_path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error while getting config %s: %s", dotenv_path, err)
		os.Exit(1)
	}

	return &Config{
		DBUrl:   os.Getenv("DB_URL"),
		APIHost: os.Getenv("API_HOST"),
		APIPort: os.Getenv("API_PORT"),
	}
}
