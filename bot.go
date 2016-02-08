package main

import (
	"github.com/tucnak/telebot"
	"log"
	"sync"
)

type bot struct {
	Connection *telebot.Bot
}

var bot_instance *bot
var bot_once sync.Once

func GetBotInstance() *bot {
	bot_once.Do(func() {
		cfg := GetConfigurationInstance()

		tb, err := telebot.NewBot(cfg.TelegramKey)
		if err != nil {
			log.Fatal(err)
		}

		bot_instance = &bot{Connection: tb}
	})

	return bot_instance
}
