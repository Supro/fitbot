package main

type BasicHandler struct {
	Strategy ReplyInterface
}

func (bh BasicHandler) Reply() {
	bh.Strategy.ReplyMessage()
}
