package main

import "fmt"

func main() {
	s := []int{9, 3, 5, 7, 8}
	var m int
	m = min(s...)
	fmt.Printf("min is : %d", m)
}

func min(paras ...int) (min int) {
	if len(paras) == 0 {
		return 0
	}
	min = paras[0]
	for _, val := range paras {
		if val < min {
			min = val
		}
	}
	return
}
