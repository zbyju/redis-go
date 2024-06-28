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

func ParseArray(input *string, len int) ([]string, error) {
	arr := []string{}
	for _ = range len {
		parsed, err := ParseDecision(input)
		if err != nil {
			return nil, err
		}
		for _, x := range parsed {
			arr = append(arr, x)
		}
	}
	return arr, nil
}

func ParseString(input *string, len int) string {
	res := (*input)[0:len]
	*input = (*input)[len:]
	return res
}

func ParseDecision(input *string) ([]string, error) {
	in := strings.Trim(*input, "\r\n \t")
	switch in[0] {
	case '*':
		split := strings.SplitN(in[1:], "\r\n", 2)
		*input = split[1]
		length, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		arr, err := ParseArray(input, length)
		if err != nil {
			return nil, err
		}
		return arr, nil
	case '$':
		split := strings.SplitN(in[1:], "\r\n", 2)
		*input = split[1]
		length, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		str := ParseString(input, length)
		return []string{str}, nil
	default:
		return nil, fmt.Errorf("unknown token to parse %c in input: %s", (*input)[0], *input)
	}
}

func ParseCommand(input string) (*Cmd, error) {
	res, err := ParseDecision(&input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res)
	if len(res) == 0 {
		return nil, fmt.Errorf("empty parsed command")
	}
	if len(res) == 1 {
		return &Cmd{Name: res[0], Args: []string{}}, nil
	}
	return &Cmd{Name: res[0], Args: res[1:]}, nil
}
