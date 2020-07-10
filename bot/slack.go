// slack api functions
package bot

import "github.com/slack-go/slack"

type Bot struct {
	apiKey *slack.Client
	status int
}

const (
	NOT_LOADING = iota
	OK
	ERROR
	LOST
)

// initalize bot
func InitBot(key string) *Bot {
	return &Bot{
		apiKey: slack.New(key),
		status: NOT_LOADING,
	}
}
