package main

import (
	"fmt"
)

/*
6 		0110
11      1011
& 按位与 0010
| 按位或 1111
^ 按位异或 相异则为1 1101
&^ 如果第二个数的对应位为1 则将第一个数的对应位改为0 0100

&& 专用于channel
||
*/
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {
	a := 1
	if a > 0 && (10/a > 1) {
		fmt.Println("ok")
	}

	fmt.Println(B, KB, MB, GB)
}
