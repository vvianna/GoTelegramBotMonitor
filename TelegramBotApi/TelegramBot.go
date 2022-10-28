package TelegramBotApi

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ExecuteBot(TELEGRAM_APITOKEN string) {
	log.Println("Starting Telegram Bot...")
	log.Println("Token : ", TELEGRAM_APITOKEN)
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_APITOKEN)
	if err != nil {
		log.Panic(err)
	}
	//bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		log.Printf("Received Message: %s by %s", update.Message.Text, update.Message.From.UserName)
		msg.Text = ConfigureMessage(update.Message)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

	}
}
