package main

import (
	"strings"
)

type helper struct {
	message map[string]string
}

//hellper add slack thread response for popular problems
func (h *helper) match(msg string) string {
	for matchWord, slackMsg := range h.message {
		if strings.Contains(msg, matchWord) {
			return slackMsg
		}
	}
	return ""
}
