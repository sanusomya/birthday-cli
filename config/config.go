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

func GetTelegramVariables() []string{
	godotenv.Load()
	bot_token := os.Getenv("bot_token")
	user_id := os.Getenv("user_id")
	chat_id := os.Getenv("chat_id")
	return []string{bot_token,user_id,chat_id}
}
