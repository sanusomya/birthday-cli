package config

import (
	"os"
	"github.com/joho/godotenv"
)

var Version = "0.1"

func GetUrl() string {
	godotenv.Load()
	url := os.Getenv("backend_url")
	return url
}
