package main

import (
	"github.com/tucnak/telebot"
	"log"
	//"strconv"
	"time"
)

type LastNewsReply struct {
	Message telebot.Message
}

func (lnr LastNewsReply) GetLastNews() []string {
	var lastNews []string

	db := GetDatabaseInstance()
	cp := CountParser{MessageText: lnr.Message.Text}

	rows, err := db.Connection.Query("SELECT id, slug, created_at FROM publications WHERE type='news' AND state='approved' ORDER BY id DESC LIMIT $1", cp.GetCount())

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var slug string
		var id int
		var created_at time.Time

		err = rows.Scan(&id, &slug, &created_at)
		if err != nil {
			log.Println(err)
		}

		url := "http://fireimp.ru/news/" + created_at.Format("2006/01/02/") + slug

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
