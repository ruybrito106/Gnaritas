package connection

import (
	"golang.org/x/net/websocket"
)

type Connection struct {
	ID		string
	Socket	*websocket.Conn
}

type ResponseStart struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	Url   string       `json:"url"`
	Self  ResponseSelf `json:"self"`
}

type ResponseSelf struct {
	Id string `json:"id"`
}
type SlackMessage struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}