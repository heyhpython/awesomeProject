package main

import "fmt"

func main() {
	var min, max int
	min, max = MinMax(78, 21)
	fmt.Printf("Min: %d  Max: %d", min, max)
}

func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}
