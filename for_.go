package main

import "fmt"

// go语言只有for循环

func main() {
	a := 1
	//for {
	//	a++
	//	if a > 3 {
	//		break
	//	}
	//	fmt.Println(a)
	//}
	for a <= 3 {
		a++
		fmt.Println(a)
	}
	fmt.Println("Over")
}
