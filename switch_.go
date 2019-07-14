package main

import "fmt"

func main() {
	// switch 语句内初始化的变量是局部变量作用于语句内
	switch a := 1;{
	case a >= 0:
		fmt.Println("a=0")
		//fallthrough
		// 继续判断一下的case 否则跳出switch
	case a > 1:
		fmt.Println("a=1")
	default:
		fmt.Println("Not match")
	}

}
