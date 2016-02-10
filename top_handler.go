package main

type TopHandler struct {
	Strategy ReplyInterface
}

func (th TopHandler) Reply() {
	th.Strategy.ReplyMessage()
}
