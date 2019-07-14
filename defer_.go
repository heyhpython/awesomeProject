package main

import "fmt"

func main() {
	for i:=0;i<3 ;i++  {
		defer fmt.Println(i)
	}
	A()
	B()
	C()
}

func A() {
	fmt.Println("func A")
}

func B() {
	defer func() {
		if e := recover(); e != nil{
			fmt.Println("recover from B")
		}
	}()
	panic("panic in B")
}

func C() {
	fmt.Println("Func c")
}