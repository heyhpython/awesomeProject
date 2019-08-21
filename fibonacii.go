package main

import (
	"fmt"
)

func main() {
	result:=0
	for i:=5000;i<=50001;i++{
		result = fib_(i)
		fmt.Println("i: ", i, "result of fib : ", result)
	}

}

func fib_(n int) (res int) {
	if n <= 1{
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}
	return
}
