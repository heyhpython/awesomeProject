package main

import "fmt"

type person struct {
	name string
	age  int
}
type human struct {
	sex int
}
type human2 struct {
	//sex int
	height float32
}
type teacher struct {
	human
	name string
}
type student struct {
	human
	human2
	name string
}

func test_strcut() {
	a := teacher{name: "zhangsan", human: human{sex: 1}}
	fmt.Println(a)
	b := student{}
	b.name = "lisi"
	b.height = 12.1
	b.sex = 2
	fmt.Println(b)

	// 结构体同名字段的优先级
	// 1.当前结构体的字段优先级高于内层结构体
	// 2.当两个内部的结构体有同名字段时会报错
}

func main() {
	a := &person{
		name: "Hawl",
		age:  25,
	}
	//a.name = "Hawl"
	//a.age = 25
	fmt.Println(a)
	test_strcut()

}
