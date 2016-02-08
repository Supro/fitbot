package main

import (
	"github.com/tucnak/telebot"
)

func HandleMessage(msg telebot.Message) {
	mh := HandlerFactory{Message: msg}
	mh.HandleMessage()
}

func main() {
	bot := GetBotInstance()
	db := GetDatabaseInstance()

	defer db.Connection.Close()

	messages := make(chan telebot.Message)
	bot.Connection.Listen(messages, 0)

	for message := range messages {
		go HandleMessage(message)
	}
}
