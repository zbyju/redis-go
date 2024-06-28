package logic

import (
	"fmt"
	"strconv"
	"strings"
)

type Cmd struct {
	Name string
	Args []string
}

func ParseArray(input *string, len int) ([]interface{}, error) {
	arr := make([]interface{}, len)
	for i := range len {
		parsed, err := ParseDecision(input)
		if err != nil {
			return nil, err
		}
		arr[i] = parsed
	}
	*input = (*input)[len:]
	return arr, nil
}

func ParseString(input *string, len int) string {
	res := (*input)[0:len]
	*input = (*input)[len:]
	return res
}

func ParseDecision(input *string) (interface{}, error) {
	in := *input
	switch in[0] {
	case '*':
		split := strings.SplitN(in[1:], "\r\n", 2)
		length, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		arr, err := ParseArray(&split[1], length)
		if err != nil {
			return nil, err
		}
		return arr, nil
	case '$':
		split := strings.SplitN(in[1:], "\r\n", 2)
		length, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		str := ParseString(&split[1], length)
		return str, nil
	default:
		return nil, fmt.Errorf("unknown token to parse %c in input: %s", input[0], input)
	}
}

func ParseCommand(input string) (Cmd, error) {
	res, err := ParseDecision(input)
	if err != nil {
		return Cmd{}, err
	}
	fmt.Println(res)
	return Cmd{}, nil
}
