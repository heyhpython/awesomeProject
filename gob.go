package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)
type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}
func main() {
	var network bytes.Buffer

	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	err := enc.Encode(P{3, 1, 5, "pythagpras"})
	if err != nil{
		fmt.Println(err)
	}
	var q Q
	err = dec.Decode(&q)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(q.Name, *q.X, *q.Y)
}
