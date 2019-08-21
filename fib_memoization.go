package main

import (
	"fmt"
	"time"
)

const LIM  = 41

var fibs[LIM] uint64
func main() {
	var res uint64 = 0
	start := time.Now()
	for i:=0; i<LIM; i++{
		res=fib(i)
		fmt.Printf("fib(%d) is: %d \n", i, res)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf(" cal take time : %s", delta)
}

func fib(i int) (res uint64) {
	if fibs[i] != 0 {
		res = fibs[i]
		return
	}
	if i <=1 {
		res = 1
	} else {
		res = fib(i-1) + fib(i-2)
	}
	fibs[i] = res
	return
}