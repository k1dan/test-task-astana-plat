package main

import (
	"bufio"
	"fmt"
	"github.com/k1dan/task2/count"
	"github.com/k1dan/task2/parse"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Errorf("Unable to read input: %v", err)
	}

	N, M := parse.NumberParser(text)

	a := make([][]uint8, N)
	for i := range a {
		a[i] = make([]uint8, M)
	}

	for i := 0; i < N; i++ {
		reader2 := bufio.NewReader(os.Stdin)
		text, err := reader2.ReadString('\n')
		numbers := strings.Replace(text, "\n", "", -1)
		if err != nil {
			fmt.Errorf("Unable to read input: %v", err)
		}
		for j := 0; j < len(numbers); j++ {
			c := numbers[j] - 48
			a[i][j] = c
		}
	}

	result := count.Counter(a)
	if result != 0 {
		fmt.Println(result)
	} else {
		fmt.Println("Impossible")
	}
}
