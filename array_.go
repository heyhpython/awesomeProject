package main

import "fmt"

func main() {
	// 数组固定长度
	//var a [2]int
	//a:=[2]int{1:2}
	a:=[2]int{1,2}
	var p *[2]int = &a

	fmt.Println(a)
	fmt.Println(p)
	// 返回数组的指针
	b := new([10]int)
	fmt.Println(b)

}
