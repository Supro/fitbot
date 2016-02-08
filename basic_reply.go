package main

import (
	"github.com/tucnak/telebot"
)

type BasicReply struct {
	Message telebot.Message
}

func (br BasicReply) ReplyMessage() {
	var text string

	text = "Бот понимает следующие команды:\n"
	text += "/last — 5 последних новостей\n"
	text += "/last 1 — Последняя новость. Можно ввести любую цифру от 1 до 10, бот поймет."

	bot := GetBotInstance()

	bot.Connection.SendMessage(br.Message.Chat, text, nil)
}
