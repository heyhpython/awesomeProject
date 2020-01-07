package main

import "fmt"

func main() {
	for i:=0;i<1000000;i++{
		go func(a int) {fmt.Println(i)}(i)
	}
}
