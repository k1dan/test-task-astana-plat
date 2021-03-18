package parse

import (
	"strconv"
	"strings"
)

func NumberParser(text string) (N int, M int) {
	numbers := strings.Replace(text, "\n", "", -1)
	n := strings.Split(numbers, " ")
	N, _ = strconv.Atoi(n[0])
	M, _ = strconv.Atoi(n[1])
	return N, M
}
