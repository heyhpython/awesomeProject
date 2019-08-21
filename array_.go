package main

import "fmt"

func main() {
	a := [5]int{1,2,3,4,5}
	f(a) // a 为数组
	fp(&a)
	b := new([5] int) // b 为数组的引用
	f(*b)
	fp(b)
}

func f(a [5]int) {
	fmt.Println(a)
}
func fp(a *[5]int) {
	fmt.Println(a)
}
