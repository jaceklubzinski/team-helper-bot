package main

import (
	"fmt"
	"strings"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

type envConfig struct {
	SlackAuthToken string `required:"true" split_words:"true"`
}

func runDB() (helper, *store) {
	var problemHelper helper
	problemHelper.message = make(map[string]string)

	dbClient, err := connectDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatalln("Can't create database connection")
	}
	db := newDB(dbClient, problemHelper)
	err = db.createTable()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatalln("Can't create helper table")
	}
	err = db.getRow()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatalln("Can't get rows from helper table")
	}
	return problemHelper, db
}

func runSlack(env envConfig) *slackClient {
	api := slack.New(
		env.SlackAuthToken,
		slack.OptionDebug(false),
	)

	rtm := api.NewRTM()
	return newSlackClient(rtm)
}

func main() {
	var env envConfig

	err := envconfig.Process("helperbot", &env)
	if err != nil {
		fmt.Println(err.Error())
	}

	problemHelper, db := runDB()

	s := runSlack(env)
	go s.slack.ManageConnection()

	commands := command{db, s, problemHelper}

	for msg := range s.slack.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			msg := ev.Msg

			if msg.SubType != "" {
				break // We're only handling normal messages.
			}
			// only accept standard channel post
			direct := strings.HasPrefix(msg.Channel, "D")
			if direct {
				continue
			}

			//bot command with mention
			if strings.Contains(msg.Text, "@"+s.slack.GetInfo().User.ID) {
				err := commands.params(msg)
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
					}).Warn("Can't execute bot command")
				}
			} else {
				//catch-all reaction to response to greetings
				if emoji := greetings(msg.Text); emoji != "" {
					err = s.reaction(msg, emoji)
					if err != nil {
						log.WithFields(log.Fields{
							"error": err,
						}).Warn("Can't add reaction")
					}
				}

				//catch-all response to popular problems
				if match := problemHelper.match(msg.Text); match != "" {
					s.simpleMsg(msg, match)
				}
			}
		case *slack.ConnectedEvent:
			log.WithFields(log.Fields{}).Info("Connected to Slack")

		case *slack.InvalidAuthEvent:
			log.WithFields(log.Fields{}).Info("Invalid token")
			return
		}
	}
}
