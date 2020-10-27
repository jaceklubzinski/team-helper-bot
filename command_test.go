package main

import (
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

type mockDB struct{}

//nolint:deadcode,unused
type MockStorer interface {
	createTable() error
	deleteAll() error
	deleteRow(title string) error
	addRow(title, desc string) error
	getRow() error
}

func (d *mockDB) createTable() error {
	return nil
}

func (d *mockDB) deleteAll() error {
	return nil
}

func (d *mockDB) deleteRow(title string) error {
	return nil
}

func (d *mockDB) addRow(title, desc string) error {
	return nil
}

func (d *mockDB) getRow() error {
	return nil
}

type mockSlackClient struct{}

//nolint:deadcode,unused
type mockSlacker interface {
	reaction(m slack.Msg, r string) error
	simpleMsg(msg slack.Msg, text string)
	postMsg(msg slack.Msg, attachment slack.Attachment) error
}

func (c *mockSlackClient) reaction(m slack.Msg, r string) error {
	return nil
}

func (c *mockSlackClient) simpleMsg(msg slack.Msg, text string) {

}

func (c *mockSlackClient) postMsg(msg slack.Msg, attachment slack.Attachment) error {
	return nil
}

func TestCommands(t *testing.T) {
	db := &mockDB{}
	slackClient := &mockSlackClient{}
	commands := command{db, slackClient}

	// help command
	msg := slack.Msg{
		Text: "@bot help",
	}
	err := commands.params(msg)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when parse help bot command", err)
	}

	// add command short
	msg = slack.Msg{
		Text: "@bot add shortTest shortDesc",
	}
	err = commands.params(msg)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when parse add short bot command", err)
	}

	// add command long
	msg = slack.Msg{
		Text: "@bot add \"Long Test\" \"Long Desc\"",
	}
	err = commands.params(msg)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when parse add short bot command", err)
	}
	assert.Equal(t, "Long Desc", hellperMessages["Long Test"])

	// add command not enough parameters
	msg = slack.Msg{
		Text: "@bot add",
	}
	err = commands.params(msg)
	if err == nil {
		t.Errorf("Expect error for input %s", msg.Text)
	}

	// del command
	hellperMessages["TestDel"] = "Delete"
	msg = slack.Msg{
		Text: "@bot del \"TestDel\"",
	}
	err = commands.params(msg)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when parse del bot command", err)
	}
	assert.NotEqual(t, "Delete", hellperMessages["TestDel"])

	// del command not enough parameters
	msg = slack.Msg{
		Text: "@bot del",
	}
	err = commands.params(msg)
	if err == nil {
		t.Errorf("Expect error for input %s", msg.Text)
	}

	// fix command
	hellperMessages["Fix1"] = "Delete"
	hellperMessages["Fix2"] = "Delete"
	msg = slack.Msg{
		Text: "@bot fix",
	}
	err = commands.params(msg)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when parse del bot command", err)
	}
	assert.NotEqual(t, "Delete", hellperMessages["fix1"])
	assert.NotEqual(t, "Delete", hellperMessages["fix2"])
}
