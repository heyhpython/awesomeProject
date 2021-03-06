package main

import (
	"fmt"
	"sort"
)

func main() {
	// map创建
	//var m map[int]string
	m := make(map[int]string)
	m[1] = "ok"
	m[2] = "hah"
	fmt.Println(m)
	//delete(m,1 )
	for k, v := range m {
		fmt.Println(k, v)
	}
	mapSort()
	fmt.Println("+++++")
	mapReverse()
	map_slice()
}

func mapSort() {
	m := map[int]string{3: "a", 2: "b", 1: "c"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)
}

func mapReverse() {
	m1 := map[int]string{1: "a", 2: "b"}
	m2 := map[string]int{}
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m1)
	fmt.Println(m2)
}

func map_slice(){
	items := make([]map[int]int, 5)
	for ix:=range items{
		items[ix] = make(map[int]int, 1)
		items[ix][1] = 2
	}
	fmt.Printf("version A: value of items :%v \n", items)

	items2 := make([]map[int]int, 5)
	for _,item := range items2{
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Printf("version B: value od items2 :%v \n", items2)
}