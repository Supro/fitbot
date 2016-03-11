package main

import (
	"github.com/tucnak/telebot"
)

type HelpReply struct {
	Message telebot.Message
}

func (hr HelpReply) ReplyMessage() {
	var text string

	text = "Бот понимает следующие команды:\n"
	text += "/last — 5 последних новостей\n"
	text += "/last 1 — Последняя новость. Можно ввести любую цифру от 1 до 10, бот поймет.\n"
	//text += "/top — 5 самых популярных новостей"

	bot := GetBotInstance()

	bot.Connection.SendMessage(hr.Message.Chat, text, nil)
}
