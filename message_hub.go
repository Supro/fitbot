package main

import (
	"github.com/tucnak/telebot"
)

type MessageHub struct {
	Message chan telebot.Message
}

func (mh *MessageHub) Run() {
	bot := GetBotInstance()
	mh.Message = make(chan telebot.Message)
	bot.Connection.Listen(mh.Message, 0)

	for {
		message := <-mh.Message
		mh := HandlerFactory{Message: message}
		mh.HandleMessage()
	}
}
