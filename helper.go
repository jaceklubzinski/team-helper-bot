package main

import (
	"strings"
)

var hellperMessages = make(map[string]string)

//hellper add slack thread response for popular problems
func helper(msg string) string {
	for matchWord, slackMsg := range hellperMessages {
		if strings.Contains(msg, matchWord) {
			return slackMsg
		}
	}
	return ""
}
