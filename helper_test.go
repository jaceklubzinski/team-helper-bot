package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelper(t *testing.T) {
	var problemHelper helper
	title := "TestTitle"
	desc := "TestDesc"
	problemHelper.message[title] = desc

	match := problemHelper.match(title)

	assert.Equal(t, match, desc)
}
