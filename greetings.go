package main

import (
	"math/rand"
	"strings"
)

//greetings add slack reaction response to greetings
func greetings(msg string) string {
	slackGreetings := map[string]bool{
		"hej":      true,
		"hello":    true,
		"hejo":     true,
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
		"cze":      true,
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

	greetingsWord := strings.Split(strings.ToLower(msg), " ")

	if slackGreetings[greetingsWord[0]] {
		i := rand.Intn(len(slackGreetingsEmoji))
		return slackGreetingsEmoji[i]
	}
	return ""
}
