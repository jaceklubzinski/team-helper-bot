package main

import (
	"strings"

	"github.com/slack-go/slack"
)

//slackClient to manage slack connections
type slackClient struct {
	slack *slack.RTM
}

type slacker interface {
	reaction(m slack.Msg, r string) error
	simpleMsg(msg slack.Msg, text string)
	postMsg(msg slack.Msg, attachment slack.Attachment) error
}

//newSlackClient constructor for new client
func newSlackClient(client *slack.RTM) *slackClient {
	return &slackClient{slack: client}
}

func (s *slackClient) reaction(m slack.Msg, r string) error {
	// Grab a reference to the message.
	msgRef := slack.NewRefToMessage(m.Channel, m.Timestamp)
	if err := s.slack.AddReaction(r, msgRef); err != nil {
		return err
	}
	return nil
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
