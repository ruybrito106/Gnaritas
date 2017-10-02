package slack_config

import (
	"os"
)

const (
	RTM_URL        = "https://slack.com/api/rtm.start"
	SLACK_API_ROOT = "https://api.slack.com/"
)

var (
	TOKEN             = os.Getenv("TOKEN")
	ID_COUNTER uint64 = 0
)
