package main

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s1 = s1[1:3]
	println(len(s1[1:1]))
	println(s1)
	s1 = s1[:len(s1)+2]
	println(s1)
}
