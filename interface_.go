package main

import (
	"fmt"
	"math"
	"reflect"
)

type Square struct {
	side float32
}

type Circle struct {
	side float32
}

type Shaper interface {
	Area() float32
}

func (self *Square)Area() float32  {
	return self.side * self.side
}
func (self *Circle)Area() float32  {

	return self.side * self.side * math.Pi
}
func main() {
	//var areaIntf Shaper
	sq := new(Square)
	sq.side = 5
	areaIntf := Shaper(sq)
	if t, ok := areaIntf.(*Square); ok{
		fmt.Printf("the type of areaIntf is : %T\n", t)
	}
	ci := new(Circle)
	ci.side = 5
	areaIntf = Shaper(ci)
	fmt.Println(reflect.TypeOf(areaIntf))
	if u, ok := areaIntf.(*Circle); ok{
		fmt.Printf("the type of areaIntf is : %T\n", u)
	}else {
		//fmt.Printf("the type of areaIntf is : %T\n", u)
	}
}
