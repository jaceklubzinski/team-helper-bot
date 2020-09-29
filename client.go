package main

import (
	"strings"

	"github.com/slack-go/slack"
)

//slackClient to manage slack connections
type slackClient struct {
	slack *slack.RTM
	db    *store
}

//newSlackClient constructor for new client
func newSlackClient(client *slack.RTM, db *store) *slackClient {
	return &slackClient{slack: client, db: db}
}

func (s *slackClient) simpleMsg(msg slack.Msg, text string) {
	// Create a response object.
	resp := s.slack.NewOutgoingMessage(text, msg.Channel)
	// Respond in thread if not a direct message.
	if !strings.HasPrefix(msg.Channel, "D") {
		resp.ThreadTimestamp = msg.Timestamp
	}

	// Respond in same thread if message came from a thread.
	if msg.ThreadTimestamp != "" {
		resp.ThreadTimestamp = msg.ThreadTimestamp
	}
	s.slack.SendMessage(resp)
}

func (s *slackClient) postMsg(msg slack.Msg, attachment slack.Attachment) error {
	//reponse in thread new one or existing
	var ThreadTimestamp string
	if msg.ThreadTimestamp != "" {
		ThreadTimestamp = msg.ThreadTimestamp
	} else {
		ThreadTimestamp = msg.Timestamp
	}
	_, _, err := s.slack.PostMessage(
		msg.Channel,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionTS(ThreadTimestamp),
		slack.MsgOptionAsUser(true),
		slack.MsgOptionUser(s.slack.GetInfo().User.ID),
	)
	if err != nil {
		return err
	}
	return nil
}
