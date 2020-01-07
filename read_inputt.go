package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter something")
	for {
		input, err = inputReader.ReadString('\n')
		if err == nil {
			fmt.Printf("The input is :%b", input)
		}
	}

}
