package main

import "fmt"

func main() {
	var reply int
	Multiply(10, 5, &reply)
	fmt.Println("reply is :", reply)
}

func Multiply(a int, b int, reply *int) {
	*reply = a * b

}
