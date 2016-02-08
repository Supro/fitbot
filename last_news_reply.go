package main

import (
	"github.com/tucnak/telebot"
	"log"
	"strconv"
)

type LastNewsReply struct {
	Message telebot.Message
}

func (lnr LastNewsReply) GetLastNews() []string {
	var lastNews []string

	db := GetDatabaseInstance()
	cp := CountParser{MessageText: lnr.Message.Text}

	rows, err := db.Connection.Query("SELECT id, slug FROM items WHERE type='news' ORDER BY id DESC LIMIT $1", cp.GetCount())

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var slug string
		var id int

		err = rows.Scan(&id, &slug)
		if err != nil {
			log.Println(err)
		}

		url := "http://fireimp.ru/news/" + slug + "-" + strconv.Itoa(id+100)

		lastNews = append(lastNews, url)
	}

	return lastNews
}

func (lnr LastNewsReply) ReplyMessage() {
	lastNews := lnr.GetLastNews()

	bot := GetBotInstance()

	for _, url := range lastNews {
		bot.Connection.SendMessage(lnr.Message.Chat, url, nil)
	}
}
