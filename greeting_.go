package main

import (
	"awesomeProject/greeting"
	"fmt"
)

func main() {
	if greeting.Day(){
		fmt.Println("Good Day")
	} else if greeting.Night(){
		fmt.Println("Good Night")
	} else {
		fmt.Println("misery")
	}
}
