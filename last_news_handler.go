package main

type LastNewsHandler struct {
	Strategy ReplyInterface
}

func (lnh LastNewsHandler) Reply() {
	lnh.Strategy.ReplyMessage()
}
