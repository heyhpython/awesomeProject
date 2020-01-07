package main

import (
	"flag"
	"fmt"
	"runtime"
)

var ngoroutine = flag.Int("n", 10, "how many goroutines")

func f_(left, right chan int) {
	r := <- right
	fmt.Printf("r:%d\n" ,r)
	left <- 1 + r
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f_(left, right)
	}
	right <- 0      // bang!
	runtime.NumCPU()
	x := <-leftmost // wait for completion
	fmt.Println(x)  // 100000, ongeveer 1,5 s
}