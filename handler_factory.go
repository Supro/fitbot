package main

import (
	"github.com/tucnak/telebot"
	"regexp"
)

type HandlerFactoryInterface interface {
	HandleMessage()
	CreateHandler() HandlerInterface
}

type HandlerFactory struct {
	Message telebot.Message
}

func (hf HandlerFactory) HandleMessage() {
	hi := hf.CreateHandler()

	hi.Reply()
}

func (hf HandlerFactory) CreateHandler() HandlerInterface {
	lastNewsRegexp, _ := regexp.Compile("^/last( (1|2|3|4|5|6|7|8|9|10))?$")
	helpRegexp, _ := regexp.Compile("^/help$")
	//topRegexp, _ := regexp.Compile("^/top$")

	var handler HandlerInterface

	if lastNewsRegexp.MatchString(hf.Message.Text) {
		handler = LastNewsHandler{Strategy: LastNewsReply{Message: hf.Message}}
	} else if helpRegexp.MatchString(hf.Message.Text) {
		handler = HelpHandler{Strategy: HelpReply{Message: hf.Message}}
		//} else if topRegexp.MatchString(hf.Message.Text) {
		//	handler = TopHandler{Strategy: TopReply{Message: hf.Message}}
	} else {
		handler = BasicHandler{Strategy: BasicReply{Message: hf.Message}}
	}

	return handler
}
