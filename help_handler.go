package main

type HelpHandler struct {
	Strategy ReplyInterface
}

func (hh HelpHandler) Reply() {
	hh.Strategy.ReplyMessage()
}
