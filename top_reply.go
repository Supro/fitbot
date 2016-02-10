package main

import (
	"github.com/tucnak/telebot"
	"log"
	"strconv"
	"time"
)

type TopReply struct {
	Message telebot.Message
}

func (tr TopReply) GetLastNews() []string {
	var topNews []string

	db := GetDatabaseInstance()

	t := time.Now()
	ts := t.Format("2006-01-02")
	ts += " 00:00:00 +0000"

	rows, err := db.Connection.Query("SELECT items.id, slug FROM items LEFT OUTER JOIN item_statistics ON items.id = item_statistics.item_id WHERE item_statistics.updated_at >= $1 AND items.type='news' ORDER BY item_statistics.today DESC LIMIT 5", ts)

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

		topNews = append(topNews, url)
	}

	return topNews
}

func (tr TopReply) ReplyMessage() {
	topNews := tr.GetLastNews()

	bot := GetBotInstance()

	for _, url := range topNews {
		bot.Connection.SendMessage(tr.Message.Chat, url, nil)
	}
}
