package slack_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync/atomic"

	logging "github.com/op/go-logging"

	"github.com/ubeedev/Gnaritas/src/connection"
	"github.com/ubeedev/Gnaritas/src/slack_config"

	"golang.org/x/net/websocket"
)

type SlackService struct {
	logger *logging.Logger
}

func NewSlackService(logger *logging.Logger) SlackService {
	var svc SlackService
	svc = SlackService{logger}
	return svc
}

func (s SlackService) SlackStartConnection() (*connection.Connection, error) {

	slackURL := fmt.Sprintf("%s?token=%s", slack_config.RTM_URL, slack_config.TOKEN)
	response, err := http.Get(slackURL)

	if err != nil {
		return nil, err
	}

	if int(response.StatusCode) != 200 {
		return nil, fmt.Errorf("GET request do slack API failed with response %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		return nil, err
	}

	var responseObject connection.ResponseStart
	err = json.Unmarshal(body, &responseObject)

	if err != nil {
		return nil, err
	}

	if !responseObject.Ok {
		return nil, fmt.Errorf("Slack error: %s", responseObject.Error)
	}

	conn, err := websocket.Dial(responseObject.Url, "", slack_config.SLACK_API_ROOT)

	if err != nil {
		return nil, err
	}

	return &connection.Connection{
		ID:     responseObject.Self.Id,
		Socket: conn,
	}, nil
}

func (s SlackService) SlackStopConnection(conn *connection.Connection) error {
	defer conn.Socket.Close()
	return nil
}

func (s SlackService) SlackReceiveMessage(conn *connection.Connection) (*connection.SlackMessage, error) {
	msg := connection.SlackMessage{}
	err := websocket.JSON.Receive(conn.Socket, &msg)

	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (s SlackService) SlackSendMessage(conn *connection.Connection, message *connection.SlackMessage) (*connection.SlackMessage, error) {
	message.Id = atomic.AddUint64(&slack_config.ID_COUNTER, 1)
	err := websocket.JSON.Send(conn.Socket, message)

	if err != nil {
		return nil, err
	}

	return message, nil
}
