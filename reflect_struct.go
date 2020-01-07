package main

import (
	"fmt"
	"reflect"
)

type NotKnown struct {
	s1, s2, s3 string
}

func (n NotKnown) String()string{

	//return n.s1 + "$" +  n.s2 + "$" + n.s3
	return "098"
}

func (n NotKnown) repr ()int{
	return len(n.String())
}

var s interface{} = NotKnown{"1", "2", "3"}

func main() {
	value := reflect.ValueOf(s) /*此处为何调用了String方法*/
	typ := reflect.TypeOf(s)
	fmt.Println("value: ", value)
	fmt.Println("type: ", typ)
	fmt.Println("kind: ", value.Kind())

	for i:=0; i<value.NumField(); i++{
		fmt.Printf("Field %d:%v\n", i, value.Field(i))
	}
	res := value.Method(0).Type()
	fmt.Println(res)
}