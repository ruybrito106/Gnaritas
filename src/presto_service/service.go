package presto_service

import (
	"fmt"

	presto "github.com/colinmarc/go-presto"
	logging "github.com/op/go-logging"
)

type PrestoService struct {
	logger *logging.Logger
}

func NewPrestoService(logger *logging.Logger) PrestoService {
	var svc PrestoService
	svc = PrestoService{logger}
	return svc
}

func (s PrestoService) MakeQuery(host, user, source, catalog, schema, query string) (string, error) {

	prestoQuery, err := presto.NewQuery(host, user, source, catalog, schema, query)

	if err != nil {
		return "", fmt.Errorf("Unable to query due to error: %s", err)
	}

	stringResponse := ""

	for row, _ := prestoQuery.Next(); row != nil; {
		stringRow := ""
		for _, val := range row {
			stringRow = stringRow + val.(string) + " "
		}
		stringResponse = stringResponse + stringRow + "\n"
	}

	return stringResponse, nil
}
