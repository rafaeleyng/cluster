package business

import (
	"github.com/rafaeleyng/my-remote/models"
)

var InputSlackStatusOnline = models.Input{
	Name:        "Slack Online",
	Description: "Set Slack status to 'Online'",
	Data: map[string]string{
		"text":     "",
		"emoji":    ":shipit:",
		"presence": "auto",
	},
}

var InputSlackStatusBeRightBack = models.Input{
	Name:        "Slack Be Right Back",
	Description: "Set Slack status to 'Be Right Back'",
	Data: map[string]string{
		"text":     "Volto logo",
		"emoji":    ":hourglass_flowing_sand:",
		"presence": "auto",
	},
}

var InputSlackStatusAtLunch = models.Input{
	Name:        "Slack At Lunch",
	Description: "Set Slack status to 'At Lunch",
	Data: map[string]string{
		"text":     "Almoçando",
		"emoji":    ":spaghetti:",
		"presence": "away",
	},
}

var InputMappings = map[string]models.Input{
	"1": InputSlackStatusOnline,
	"2": InputSlackStatusBeRightBack,
	"3": InputSlackStatusAtLunch,
}
