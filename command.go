package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/slack-go/slack"
)

func getRandomColor() string {
	letters := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	color := "#"
	for i := 0; i < 6; i++ {
		color += letters[rand.Intn(len(letters))]
	}
	return color
}

//command manage bot reponse for slack command
func (s *slackClient) command(msg slack.Msg) error {
	botCommand := map[string]string{
		"add":  "add problem with possible solution\n examples\nshort : `@bot add sources.LazyJDBCSource http://github.com`\n long: `@bot add \"ProxySQL Error: Access denied for user\" \"recreate container\"`",
		"list": "list all problems and solutions",
		"fix":  "Delete all problems",
		"help": "Help message",
	}

	if strings.Contains(msg.Text, "@"+s.slack.GetInfo().User.ID) {
		commandParameters := strings.Split(msg.Text, " ")
		commandParametersLong := strings.Split(msg.Text, "\"")
		switch commandParameters[1] {
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

			err := s.postMsg(msg, attachment)
			if err != nil {
				return err
			}
		case "add":
			var description string

			//default index
			titleI := 2
			descriptionI := 3
			if len(commandParametersLong) > 2 {
				commandParameters = commandParametersLong
				titleI = 1
				descriptionI = 2
			}

			if len(commandParameters) < 4 {
				s.simpleMsg(msg, ":niedobrze: Not enough number of parameters. Try `help` command")
				return errors.New("Not enough number of parameter")
			}
			title := commandParameters[titleI]
			if len(commandParameters) > 2 {
				for i := descriptionI; i < len(commandParameters); i++ {
					description = description + " " + commandParameters[i]
				}
			}
			fmt.Println("Title: ", title)
			fmt.Println("Description: ", description)
			if _, ok := hellperMessages[title]; !ok {
				err := s.db.addRow(title, description)
				if err != nil {
					return err
				}
				hellperMessages[title] = description
				s.simpleMsg(msg, ":thumbsup: Thanks for support. This problem will not bother anymore!")
			}
		case "list":

			fields := make([]slack.AttachmentField, 0)
			err := s.db.getRow()
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
			err = s.postMsg(msg, attachment)
			if err != nil {
				return err
			}

		case "fix":
			err := s.db.deleteAll()
			if err != nil {
				return err
			}
			for k := range hellperMessages {
				delete(hellperMessages, k)
			}
			s.simpleMsg(msg, ":thumbsup: All problems fixed")

		default:
			s.simpleMsg(msg, "Try `help` command")
		}
	}
	return nil
}
