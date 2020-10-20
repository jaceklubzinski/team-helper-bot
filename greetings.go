package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/slack-go/slack"
)

//greetings add slack reaction response to greetings
func (s *slackClient) greetings(msg slack.Msg) error {
	slackGreetings := map[string]bool{
		"hej":      true,
		"hello":    true,
		"witam":    true,
		"siema":    true,
		"siemka":   true,
		"siemanko": true,
		"siemano":  true,
		"bonjorno": true,
		"ahoj":     true,
		"joł":      true,
		"howgh":    true,
		"czołem":   true,
		"czesc":    true,
		"cześć":    true,
		"elo":      true,
	}

	slackGreetingsEmoji := []string{
		"howdy",
		"hellohello",
		"hellohello2",
		"hellohello3",
		"hellohello4",
		"hellohelloleft",
		"hello_there",
		"vi-tam",
		"uszanowanko",
		"ship",
		"spock-hand",
		"hand",
		"handshake",
		"raising_hand",
		"man-raising-hand",
		"bonzur",
		"good_moaning",
		"rocker",
		"the_horns",
		"hey-girl",
		"zheyguys",
		"hotdog_hello",
		"pig_hello",
		"pig_hello2",
		"pig_hello_door",
		"pikachu-hello",
	}

	greetingsWord := strings.Split(strings.ToLower(msg.Text), " ")

	if slackGreetings[greetingsWord[0]] {
		i := rand.Intn(len(slackGreetingsEmoji))
		// Grab a reference to the message.
		msgRef := slack.NewRefToMessage(msg.Channel, msg.Timestamp)
		if err := s.slack.AddReaction(slackGreetingsEmoji[i], msgRef); err != nil {
			fmt.Printf("Error adding reaction: %s", err)
			return err
		}
	}
	return nil
}
