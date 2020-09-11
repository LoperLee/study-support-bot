// slack api functions
package bot

import (
	"fmt"

	"github.com/slack-go/slack"
)

// save bot token struct
type Bot struct {
	api    *slack.Client
	status int
}

type Message struct {
	Title  string
	Attach slack.Attachment
}

// MessageType
const (
	Commit = iota + 1
	Booking
)

const (
	NOT_LOADING = iota
	OK
	ERROR
	LOST
)

// initalize bot
func InitBot(key string) *Bot {
	return &Bot{
		api:    slack.New(key),
		status: NOT_LOADING,
	}
}

// message to channel
func (bot *Bot) SendingToChannel(message Message, msgType int) {
	var channel, title string
	var attach slack.Attachment
	switch msgType {
	case Commit:
		channel = ""
		title, attach = message.Build()
		break
	default:
		panic("This is not allowed typed message.")
	}
	_, timestamp, err := bot.api.PostMessage(channel, slack.MsgOptionText(title, false), slack.MsgOptionAttachments(attach))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("[%s]Message sending complete!\n", timestamp)
}

// build message
func (com Message) Build() (string, slack.Attachment) {
	return com.Title, com.Attach
}

// Setting message
func (com Message) SetMessage(message string) {
	com.Title = message
	com.Attach = slack.Attachment{
		Color:  "#FFFFFF",
		Title:  "example attach",
		Footer: "report by bot",
	}
}
