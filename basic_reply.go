package main

import (
	"github.com/tucnak/telebot"
)

type BasicReply struct {
	Message telebot.Message
}

func (br BasicReply) ReplyMessage() {
	var text string

	text = "Бот не знает таких команд.\n"
	text += "Наберите /help\n"

	bot := GetBotInstance()

	bot.Connection.SendMessage(br.Message.Chat, text, nil)
}
