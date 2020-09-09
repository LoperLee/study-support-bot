package main

import (
	"github.com/LoperLee/study-support-bot/bot"
	"github.com/LoperLee/study-support-bot/config"
)

func main() {
	// Load configration
	config := config.LoadConfigration("config.yml")

	// init bot
	bot.InitBot(config.Slack.Token)
}
