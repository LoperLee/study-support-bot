// slack api functions
package bot

import "github.com/slack-go/slack"

// save bot token struct
type Bot struct {
	apiKey *slack.Client
	status int
}

// all message builder interface
type Message interface {
	Build() (string, slack.Attachment)
	SetMessage()
	SetAttach()
}
type Commit struct {
	Message string
	Attach  slack.Attachment
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

// message to channel
func (bot *Bot) SendingToChannel(message Message) {

}

// build Check Message
func (com Commit) SetMessage(message string) {
	com.Message = message
}
