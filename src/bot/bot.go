package bot

import (
	"github.com/ubeedev/Gnaritas/src/connection"
	presto "github.com/ubeedev/Gnaritas/src/presto_service"
	slack "github.com/ubeedev/Gnaritas/src/slack_service"
)

type GnaritasBot struct {
	Connection    *connection.Connection
	SlackService  slack.SlackService
	PrestoService presto.PrestoService
	Alive         bool
}
