package main

import (
	"fmt"
	"time"
)

func longWait() {
	fmt.Println(" begin long wait")
	time.Sleep(5)
	fmt.Println(" end long wait")
}

func shortWait() {
	fmt.Println(" begin short wait")
	time.Sleep(2)
	fmt.Println(" end short wait")
}

func main() {
	fmt.Println("in main")
	go longWait()
	go shortWait()
	fmt.Println("abount to sleep in main")
	time.Sleep(2)
	fmt.Println("end main")
}