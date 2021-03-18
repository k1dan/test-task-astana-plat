package parser

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errParse = errors.New("not enough arguments")
)

func Parse(text string) (operation string, parsedText string, shift int32, err error) {
	splitted := strings.Split(text, " ")
	start := 0
	end := 0
	for i, value := range splitted {
		if value == "-t" {
			start = i + 1
		}
		if value == "-k" {
			end = i
		}
	}
	operation = splitted[0]
	for i := start; i < end; i++ {
		parsedText = parsedText + splitted[i]
		if i != end-1 {
			parsedText = parsedText + " "
		}
	}

	if start == 0 || end == 0 {
		return "", "", 0, errParse
	}
	num := splitted[end+1]
	n := strings.Replace(num, "\n", "", -1)

	i, err := strconv.Atoi(n)
	if err != nil {
		return "", "", 0, err
	}
	v := int32(i)
	if v < 0 || v > 26  {
		return "", "", 0, errParse
	}

	return operation, parsedText, v, nil
}
