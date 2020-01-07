package main

import "fmt"

func main() {
	for i := 0; i < 1; i++ {
		fmt.Println("value of i is :", i)
		fmt.Println("&: ", &i)
	}
	a := 123
	test(&a)
}

func test(a *int) {
	fmt.Println("a: ", a)
	fmt.Println("&: ", &a)
	fmt.Println("*: ", *a)

}
