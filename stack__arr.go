package main

import "fmt"

type stack struct {
	data [4] int
	p1 bool
	p2 bool
	p3 bool
	p4 bool
}

func (s *stack) pop() (res int) {
	if s.p4 {
		res = s.data[3]
		s.data[3] = 0
		s.p4 = false
	} else if s.p3{
		res = s.data[2]
		s.data[2] = 0
		s.p3 = false
	}else if s.p2{
		res = s.data[1]
		s.data[1] = 0
		s.p2 = false
	}else if s.p1{
		res = s.data[0]
		s.data[0] = 0
		s.p1 = false
	}else {
		res = 0
	}

	return
}

func (s *stack)push(i int) (res bool) {
	res = true
	if !s.p1 {
		s.data[0] = i
		s.p1 = true
	} else if !s.p2{
		s.data[1] = i
		s.p2 = true
	}else if !s.p3{
		s.data[2] = i
		s.p3 = true
	}else if !s.p4{
		s.data[3] = i
		s.p4 = true
	}else {
		res = false
	}
	return
}

func main() {
	var s stack
	for i:=0;i<4;i++{
		s.push(i)
		fmt.Println(s)
	}
	for i:=0;i<4;i++{
		fmt.Println(s)
		s.pop()

	}

}