package main

import "fmt"

func main() {
	fmt.Println("enter main")
	defer fmt.Println("leave main")
	hah()
}

func trace(a string) string {
	fmt.Println("entering: ", a)
	return a
}

func untrace(a string) {
	fmt.Println("leaving: ", a)
}

func hah() {
	//trace("haha")
	//defer untrace("haha")
	defer untrace(trace("hah"))
	fmt.Println("func haha")
}
