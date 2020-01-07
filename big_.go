package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1993)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)  // ip= im * in
	fmt.Println(im, in ,ip)

	rm := big.NewRat(math.MaxInt64, 1993)
	fmt.Println(rm)

}