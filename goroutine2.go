package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	ch <- "hello"
	ch <- "world"
}

func getData(ch chan string)  {
	var input string
	for{
		input = <- ch
		fmt.Println(input)
	}
}

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}