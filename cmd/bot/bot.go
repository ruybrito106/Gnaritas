package main

import (
	"github.com/ubeedev/Gnaritas/src/bot_service"
	"github.com/ubeedev/Gnaritas/src/logger"
	"github.com/ubeedev/Gnaritas/src/presto_service"
	"github.com/ubeedev/Gnaritas/src/slack_service"
)

func main() {

	logger := logger.NewLogger()

	slackService := slack_service.NewSlackService(logger)
	prestoService := presto_service.NewPrestoService(logger)
	botService := bot_service.NewService(logger)

	bot, err := botService.MakeNewGnaritasBot(slackService, prestoService)

	if err != nil {
		logger.Error(err)
		return
	}

	for {
		botService.BotStartListening(bot)
	}

}
