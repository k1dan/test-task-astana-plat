package main

import (
	"bufio"
	"fmt"
	"github.com/k1dan/task1/encryption"
	"github.com/k1dan/task1/parser"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter command and arguments")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Errorf("Unable to read input: %v", err)
	}
	operation, parsedText, shift, err := parser.Parse(text)
	if err != nil {
		fmt.Errorf("Unable to parse: %v", err)
	}
	if operation == "encrypt" {
		encr := encryption.Encrypt(parsedText, shift)
		fmt.Println(encr)
	} else if operation == "decrypt" {
		decr := encryption.Dencrypt(parsedText, shift)
		fmt.Println(decr)
	} else {
		fmt.Println("Wrong command!")
	}
}

