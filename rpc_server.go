package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}

func main() {
	cals := new(Args)
	rpc.Register(cals)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", "localhost:8888")  // 创建套接字
	if err != nil{
		fmt.Println(err.Error())
	}
	go http.Serve(listener, nil)  // 在套接字上运行服务
	time.Sleep(1000e9)

}
