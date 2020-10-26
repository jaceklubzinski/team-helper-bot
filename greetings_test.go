package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetings(t *testing.T) {
	greet := "hello , what's up"
	match := greetings(greet)
	assert.NotEmpty(t, match)

	greetEmpty := "what's up"
	match = greetings(greetEmpty)
	assert.Empty(t, match)
}
