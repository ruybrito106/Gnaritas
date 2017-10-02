package util

import "strings"

func ValidMessage(msg string) bool {
	parts := strings.Fields(msg)

	var hasH, hasU, hasSRC, hasQRY bool
	for _, part := range parts {
		switch part {
		case "-h":
			hasH = true
		case "-u":
			hasU = true
		case "-src":
			hasSRC = true
		case "-qry":
			hasQRY = true
		default:
		}
	}

	return hasH && hasU && hasSRC && hasQRY
}

func ParseQuery(msg string) (string, string, string, string, string, string) {

	var user, host, source, catalog, schema, query string

	parts := strings.Split(msg, "-")

	var lastCharIndex int

	for index, part := range parts {

		if index == len(parts)-1 {
			lastCharIndex = len(part)
		} else {
			lastCharIndex = len(part) - 1
		}

		if strings.HasPrefix(part, "h") {
			host = part[2:lastCharIndex]
		} else if strings.HasPrefix(part, "u") {
			user = part[2:lastCharIndex]
		} else if strings.HasPrefix(part, "src") {
			source = part[4:lastCharIndex]
		} else if strings.HasPrefix(part, "qry") {
			query = part[4:lastCharIndex]
		} else if strings.HasPrefix(part, "c") {
			catalog = part[2:lastCharIndex]
		} else if strings.HasPrefix(part, "sch") {
			schema = part[4:lastCharIndex]
		}
	}

	return host, user, source, catalog, schema, query

}
