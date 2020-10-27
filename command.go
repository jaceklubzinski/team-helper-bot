package main

import (
	"errors"
	"math/rand"
	"strings"

	"github.com/slack-go/slack"
)

type command struct {
	db    storer
	slack slacker
}

func getRandomColor() string {
	letters := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	color := "#"
	for i := 0; i < 6; i++ {
		color += letters[rand.Intn(len(letters))]
	}
	return color
}

//params manage bot reponse for slack params
func (c *command) params(msg slack.Msg) error {
	botCommand := map[string]string{
		"add":     "add problem with possible solution\n examples\nshort : `@bot add sources.LazyJDBCSource http://github.com`\n long: `@bot add \"ProxySQL Error: Access denied for user\" \"recreate container\"`",
		"del":     "del single row\n example: `@bot del \"single row\"`",
		"list":    "list all problems and solutions",
		"fix-all": "Delete all problems",
		"help":    "Help message",
	}

	commandParam := strings.Split(msg.Text, " ")
	commandParamLong := strings.Split(msg.Text, "\"")
	switch commandParam[1] {
	case "help":
		fields := make([]slack.AttachmentField, 0)

		for k, v := range botCommand {
			fields = append(fields, slack.AttachmentField{
				Title: k,
				Value: v,
			})
		}

		attachment := slack.Attachment{
			Pretext: "Bot command list",
			Color:   getRandomColor(),
			Fields:  fields,
		}

		err := c.slack.postMsg(msg, attachment)
		if err != nil {
			return err
		}
	case "add":
		var description string

		titleI := 2       //default index for value
		descriptionI := 3 //default index for description

		//if command params are using " " use long format
		if len(commandParamLong) > 2 {
			commandParam = commandParamLong
			titleI = 1       //long param default index for value
			descriptionI = 3 //long param default index for description
		}

		if len(commandParam) < 4 {
			c.slack.simpleMsg(msg, ":niedobrze: Not enough number of parameters. Try `help` command")
			return errors.New("Not enough number of parameter")
		}

		title := commandParam[titleI]
		description = commandParam[descriptionI]

		if _, ok := hellperMessages[title]; !ok {
			err := c.db.addRow(title, description)
			if err != nil {
				return err
			}
			hellperMessages[title] = description
			c.slack.simpleMsg(msg, ":thumbsup: Thanks for support. This problem will not bother anymore!")
		}

	case "del":
		if len(commandParamLong) < 2 {
			c.slack.simpleMsg(msg, ":niedobrze: Not enough number of parameters. Try `help` command")
			return errors.New("Not enough number of parameter")
		}

		title := commandParamLong[1]

		if _, ok := hellperMessages[title]; ok {
			err := c.db.deleteRow(title)
			if err != nil {
				return err
			}
			delete(hellperMessages, title)
			c.slack.simpleMsg(msg, ":thumbsup: Problem fixed")
		}
	case "list":

		fields := make([]slack.AttachmentField, 0)
		err := c.db.getRow()
		if err != nil {
			return err
		}
		for k, v := range hellperMessages {
			fields = append(fields, slack.AttachmentField{
				Title: k,
				Value: v,
			})
		}

		attachment := slack.Attachment{
			Pretext: "Matching words",
			Color:   getRandomColor(),
			Fields:  fields,
		}
		err = c.slack.postMsg(msg, attachment)
		if err != nil {
			return err
		}

	case "fix-all":
		err := c.db.deleteAll()
		if err != nil {
			return err
		}
		for k := range hellperMessages {
			delete(hellperMessages, k)
		}
		c.slack.simpleMsg(msg, ":thumbsup: All problems fixed")

	default:
		c.slack.simpleMsg(msg, "Try `help` command")
	}
	return nil
}
