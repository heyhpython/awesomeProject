package main

import "fmt"

// 常量

const a int = 1
const (
	b = 2
	c = "qqwewqeqd"
	d
	e = len(c)
)
const (
	z = iota
	x
	v
	n
	// iota与声明的顺序有关系 与其出现的次数无关 组内定义一个常量 iota加一
)

func main() {
	fmt.Println(a, b, c, d, e)
	fmt.Println(z, x, v, n)
}
