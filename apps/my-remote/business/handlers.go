package business

import (
	"os"

	"github.com/nlopes/slack"
	"github.com/rafaeleyng/my-remote/models"
)

func handleSlackStatus(data map[string]string) error {
	slackToken := os.Getenv("SLACK_TOKEN")
	var api = slack.New(slackToken)
	// TODO handle errors
	api.SetUserCustomStatus(data["text"], data["emoji"])
	api.SetUserPresence(data["presence"])
	return nil
}

func HandleInput(input models.Input) {
	switch input.Name {
	case InputSlackStatusOnline.Name:
		{
			handleSlackStatus(InputSlackStatusOnline.Data)
		}
	case InputSlackStatusBeRightBack.Name:
		{
			handleSlackStatus(InputSlackStatusBeRightBack.Data)
		}
	case InputSlackStatusAtLunch.Name:
		{
			handleSlackStatus(InputSlackStatusAtLunch.Data)
		}
	}
}
