package main

import "fmt"

type Student struct {
	name string
	age int
}

func main() {
	s := Student{
		name:"heyuhao", age:18,
	}

	printStudent(s)
}

func printStudent(s Student) {
	fmt.Println(s.name)
	fmt.Println(s.age)
}