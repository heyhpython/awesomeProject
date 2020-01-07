package main

import (
	"fmt"
)

func main() {
	var s1 []int
	fmt.Println(s1)

	a := [10]int{}
	s2 := a[:5]
	fmt.Println(s2)
	s3 := make([]int, 3, 100)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
	temp()
}

func temp() {
	s1 := make([]int, 3, 6)
	fmt.Println(&s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(&s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(&s1)

}
