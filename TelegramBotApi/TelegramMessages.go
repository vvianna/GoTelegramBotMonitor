package TelegramBotApi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	VERSION               = "0.0.1"
	ERROR_COMMAND_MESSAGE = "Sorry my friend! Command handling failure... =("
	ERROR_MESSAGE         = "Sorry my friend! Message handling failure... =("
	COMMAND_ABOUT         = "Writed By Victor Monteiro \nVersion: " + VERSION + "\n" + "E-mail: vmoonteiro@comviva.com"
	SUPPORTED_COMMANDS    = "Supported Commands: \n" + "/ABOUT"
)

func ConfigureMessage(Message *tgbotapi.Message) string {
	if Message.IsCommand() {
		switch Message.Command() {
		case "ABOUT":
			return COMMAND_ABOUT

		}
		return SUPPORTED_COMMANDS

	}
	return ERROR_MESSAGE
}
