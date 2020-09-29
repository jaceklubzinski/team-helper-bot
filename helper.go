package main

import (
	"strings"

	"github.com/slack-go/slack"
)

var hellperMessages = make(map[string]string)

//hellper add slack thread response for popular problems
func (s *slackClient) hellper(msg slack.Msg) {
	for matchWord, slackMsg := range hellperMessages {
		if strings.Contains(msg.Text, matchWord) && !strings.Contains(msg.Text, "@"+s.slack.GetInfo().User.ID) {
			s.simpleMsg(msg, slackMsg)
		}
	}
}
