package logger

import (
	"os"

	logging "github.com/op/go-logging"
)

func NewLogger() *logging.Logger {

	logger := logging.MustGetLogger("Gnaritas")
	format := logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)

	back := logging.NewLogBackend(os.Stderr, "", 0)
	backFormatter := logging.NewBackendFormatter(back, format)

	logging.SetBackend(backFormatter)

	return logger

}
