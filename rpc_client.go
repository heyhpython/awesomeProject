package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	N, M int
}
func main_1() {
	client, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil{
		fmt.Println(err.Error())
	}
	args := &Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%d * %d = %d", args.M, args.N, reply)
}

func main_2() {
	client, err := rpc.DialHTTP("tcp",   "localhost:8888")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	args := &Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
func main() {
	//main_2()
	main_1()
}