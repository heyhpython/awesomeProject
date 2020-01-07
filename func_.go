package main

import "fmt"

func main() {
	f := closure(10)
	fmt.Println(f(11))
	fmt.Println(f(12))

}

func closure(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func timer_(f1 func(a int) int) func(int) int {
	return func(i int) int {
		return f1(a)
	}

}
