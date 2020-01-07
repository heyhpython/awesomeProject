package main

import "fmt"

// 跳转语句配合标签使用

func main() {
LABEL1:
	for j := 0; j < 2; j++ {
		fmt.Println("j =", j)
		for i := 1; i < 10; i++ {
			fmt.Println(i)
			if i > 3 {
				//break LABEL1
				// 直接跳出LABEL1
				//goto LABEL1
				//跳出到LABEL1 并重新开始循环LABEL1
				continue LABEL1
				// 继续执行LABEL1的循环
			}
		}
		fmt.Println("label1")
	}

}
