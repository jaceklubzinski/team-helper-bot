package main

import (
	"math/rand"
	"strings"
)

//greetings add slack reaction response to greetings
func greetings(msg string) string {
	slackGreetings := map[string]bool{
		"hej":               true,
		"hello":             true,
		"hejo":              true,
		"witam":             true,
		"siema":             true,
		"siemka":            true,
		"siemanko":          true,
		"siemano":           true,
		"bonjorno":          true,
		"ahoj":              true,
		"joł":               true,
		"howgh":             true,
		"czołem":            true,
		"czesc":             true,
		"cześć":             true,
		"cze":               true,
		"elo":               true,
		":uszanowanko:":     true,
		":howdy:":           true,
		":ahoj:":            true,
		":hellohello:":      true,
		":hellohello2:":     true,
		":hellohello3:":     true,
		":hellohello4:":     true,
		":hellohello5:":     true,
		":hellohello6:":     true,
		":hellohello7:":     true,
		":hellohello8:":     true,
		":hellohelloleft:":  true,
		":hellohelloright:": true,
		":hellohellowell:":  true,
		":hello_there:":     true,
		":pepe-witam:":      true,
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
		"ahoj",
		"dobre_ranko",
		"pepe-witam",
	}

	greetingsWord := strings.Split(strings.ToLower(msg), " ")

	if slackGreetings[greetingsWord[0]] {
		i := rand.Intn(len(slackGreetingsEmoji))
		return slackGreetingsEmoji[i]
	}
	return ""
}
