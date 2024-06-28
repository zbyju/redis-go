package logic

import (
	"fmt"
	"strings"
)

func (cmd Cmd) Execute() (*string, error) {
	switch strings.ToUpper(cmd.Name) {
	case "PING":
		return executePing(cmd)

	case "ECHO":
		return executeEcho(cmd)

	default:
		res := "Unknown command executed: " + cmd.Name
		return &res, nil
	}
}

func executePing(_ Cmd) (*string, error) {
	res := "PONG"
	return &res, nil
}

func executeEcho(cmd Cmd) (*string, error) {
	if len(cmd.Args) != 1 {
		return nil, fmt.Errorf("ECHO needs exactly one argument")
	}
	res := cmd.Args[0]
	return &res, nil
}
