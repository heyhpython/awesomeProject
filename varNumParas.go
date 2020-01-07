package main

import "fmt"

func main() {
	fmt.Println("min is ", Min(1, 2, 3, 4, 4, 5))
}

func Min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, value := range s {
		if value < min {
			min = value
		}
	}
	return min
}
