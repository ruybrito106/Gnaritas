package bot_service

import (
	"fmt"
	"strings"

	logging "github.com/op/go-logging"

	"github.com/ubeedev/Gnaritas/src/bot"
	"github.com/ubeedev/Gnaritas/src/connection"
	presto "github.com/ubeedev/Gnaritas/src/presto_service"
	slack "github.com/ubeedev/Gnaritas/src/slack_service"
	"github.com/ubeedev/Gnaritas/src/util"
)

type Service interface {
	MakeNewGnaritasBot(slack.SlackService, presto.PrestoService) (*bot.GnaritasBot, error)
	BotStartListening(*bot.GnaritasBot) error
	BotStopListening(*bot.GnaritasBot) error
}

type basicService struct {
	logger *logging.Logger
}

func NewService(logger *logging.Logger) Service {
	var svc Service
	svc = basicService{logger}
	return svc
}

func (s basicService) MakeNewGnaritasBot(slackService slack.SlackService, prestoService presto.PrestoService) (*bot.GnaritasBot, error) {

	conn, err := slackService.SlackStartConnection()

	if err != nil {
		return nil, err
	}

	bot := bot.GnaritasBot{
		Connection:    conn,
		SlackService:  slackService,
		PrestoService: prestoService,
		Alive:         false,
	}

	return &bot, nil
}

func (s basicService) BotStartListening(bot *bot.GnaritasBot) error {

	conn := bot.Connection
	slackService := bot.SlackService
	prestoService := bot.PrestoService

	message, err := slackService.SlackReceiveMessage(conn)

	if err != nil {
		s.logger.Error(err)
		return err
	}

	if message.Type == "message" && strings.HasPrefix(message.Text, "<@"+conn.ID+">") {

		s.logger.Debug("Message received: " + message.Text)

		if util.ValidMessage(message.Text) {

			user, host, source, query, schema, catalog := util.ParseQuery(message.Text)

			go func(message *connection.SlackMessage) {

				message.Text, err = prestoService.MakeQuery(host, user, source, catalog, schema, query)

				if err != nil {
					s.logger.Error("Erro: " + err.Error())
					return
				}

				slackService.SlackSendMessage(conn, message)

			}(message)

		} else {

			message.Text = fmt.Sprintf("Olá, InLoco!\n")
			slackService.SlackSendMessage(conn, message)

			message.Text = fmt.Sprintf("A sua query no presto não parece estar bem formatada...\n")
			slackService.SlackSendMessage(conn, message)

			message.Text = fmt.Sprintf("Me ajude a lhe ajudar amiguinho... Digite algo do tipo:\n")
			slackService.SlackSendMessage(conn, message)

			message.Text = fmt.Sprintf("```-h BIRL -u BIRL -src BIRL -c BIRL -sch BIRL -qry SELECT * FROM birl```")
			slackService.SlackSendMessage(conn, message)

			message.Text = fmt.Sprintf("-h: Host\n-u: User\n-src: Source\n-c: Catalog\n-sch: Schema\n-qry: Query\n")
			slackService.SlackSendMessage(conn, message)

		}
	}

	return nil

}

func (s basicService) BotStopListening(bot *bot.GnaritasBot) error {
	slackService := bot.SlackService
	err := slackService.SlackStopConnection(bot.Connection)

	if err != nil {
		return err
	}

	return nil
}
