package TelegramBotApi

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ExecuteBot(TELEGRAM_APITOKEN string) {
	log.Println("Starting Telegram Bot...")
	log.Println("Token : ", TELEGRAM_APITOKEN)
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_APITOKEN)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "help":
			var sb strings.Builder
			sb.WriteString("Welcome to Comviva BOT")
			sb.WriteString("Available Commands:")
			sb.WriteString("\n")
			sb.WriteString("TIM - ASA \n")
			sb.WriteString("/TIM_ASA_HEALTHCHECK")
			msg.Text = sb.String()
		case "ABOUT":
			var sb strings.Builder
			sb.WriteString("Writed By Victor Monteiro\n")
			sb.WriteString("Version 1.0\n")
			sb.WriteString("E-mail: vmoonteiro@comviva.com")
			msg.Text = sb.String()
		default:
			var sb strings.Builder
			sb.WriteString("Available Commands:")
			sb.WriteString("\n")
			sb.WriteString("TIM - ASA \n")
			sb.WriteString("/TIM_ASA_HEALTHCHECK")
			msg.Text = sb.String()
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
