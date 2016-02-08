package main

import (
	"regexp"
)

type CountParser struct {
	MessageText string
}

func (c CountParser) GetCount() string {
	numberRegexp, _ := regexp.Compile("[0-9]+$")

	count := numberRegexp.FindString(c.MessageText)

	if count == "" {
		count = "5"
	}

	return count
}
