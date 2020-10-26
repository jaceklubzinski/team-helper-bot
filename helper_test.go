package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelper(t *testing.T) {
	title := "TestTitle"
	desc := "TestDesc"
	hellperMessages[title] = desc

	match := helper(title)

	assert.Equal(t, match, desc)
}
