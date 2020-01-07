package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(f float64) (float64, error) {
	if f < 0{
		return 0, errors.New("math can not square negative num")
	}
	return math.Sqrt(f), nil
}

func main() {
	notFound := errors.New("url not found")
	fmt.Printf("error: %v\n", notFound)
	nums :=[]float64{16, -16}
	for _,num:=range nums{
		res, err := sqrt(num)
		if err != nil{
			fmt.Println(err)
		}else {
			fmt.Println(res)
		}
	}
}