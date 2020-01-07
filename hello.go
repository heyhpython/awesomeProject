// 声明包名
package main

// 导入包名
import (
	"fmt"
	// 重命名包
	op "os"
)

const PI = 3.14

var r = 10

type newType int

type goStruct struct {
}

type goInterface interface {
}

// 程序的入口 有且仅有一个
func main() {
	// 函数名小写代表是包的内部方法 不可被其他包调用
	// 大写则可以被外部调用
	fmt.Println("hello")
	fmt.Println(op.Geteuid())
}
