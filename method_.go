package main

import (
	"fmt"
	"math"
	"reflect"
)

type Point struct {
	x float64
	y float64
}

func (p Point)Abs() float64  {
	fmt.Println(reflect.TypeOf(p))
	return math.Sqrt(p.x*p.x + p.y*p.y)

}

func main() {
	var p1 Point
	p1.x = 3.0
	p1.y = 4.0
	fmt.Println(p1.Abs())
	//fmt.Println(reflect.TypeOf(p1))

	p2 := new(Point)
	p2.x = 6
	p2.y = 8
	fmt.Println(p2.Abs())
	//fmt.Println(reflect.TypeOf(p2))

}

