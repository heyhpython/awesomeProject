package main

import (
	"fmt"
	"runtime"
)

func main() {

	a := Adder(2)
	fmt.Println("result is:  ", a(3))
}

func Adder(a int) func(b int) int {
	fmt.Println(runtime.Caller(0))
	return func(b int) int {
		return a + b
	}

}
