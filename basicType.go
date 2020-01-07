package main

import (
	"fmt"
	"math"
	"strconv"
)

type (
	//byte int8
	rune int32
	文本   string
)

func main() {
	// 零值 当变量未赋值时，编译器自动赋值
	var a int
	var b string
	var c [10]int
	var d bool
	var e 文本
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(math.MaxFloat64)

	// var f int = 1
	// var f = 1
	f := 1
	//f=1
	fmt.Println(f)

	var g float32 = 52.0
	fmt.Println(g)
	h := int(g)
	fmt.Println(h)

	// i := string(h)
	// go中int 转字符串不能直接转 会返回其ascii码位
	i := strconv.Itoa(h)
	h, j := strconv.Atoi(i)
	fmt.Println(i, h, j)

}
