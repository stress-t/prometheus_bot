package main

import (
	"fmt"

	"gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	telegram *tgbotapi.BotAPI
}

const (
	// commandStart   = "/start"
	// commandStop    = "/stop"
	commandHelp    = "/help"
	commandChatsID = "/chat_id"

	// commandStatus     = "/status"
	// commandAlerts     = "/alerts"
	// commandSilences   = "/silences"
	// commandSilenceAdd = "/silence_add"
	// commandSilence    = "/silence"
	// commandSilenceDel = "/silence_del"

	// responseStart = "Hey, %s! I will now keep you up to date!\n" + commandHelp
	// responseStop  = "Alright, %s! I won't talk to you again.\n" + commandHelp
	responseHelp = `
I'm a Prometheus AlertManager Bot for Telegram. I will notify you about alerts.
Available commands:
` + commandChatsID + ` - Get current chat ID.
` + commandHelp + ` - Get help.
`
)

func (b *Bot) handleHelp(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, responseHelp)
	bot.Send(msg)
}
func (b *Bot) handleChatsID(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Chat id is '%d'", message.Chat.ID))
	bot.Send(msg)
}
