package main

import (
	"GoTelegramBotMonitor/TelegramBotApi"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	log.Println("Starting Application!!! Lets Go!")
	TELEGRAM_APITOKEN := goDotEnvVariable("TELEGRAM_APITOKEN")
	TelegramBotApi.ExecuteBot(TELEGRAM_APITOKEN)
}
