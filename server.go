package main

import (
	"fmt"
	"net"
)

func MakeServer(ip, port string) {
	listener, err := net.Listen("tcp", ip + ":" + port)
	if err != nil{
		fmt.Println("Error in make listener", err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("Error in accept", err.Error())
			return
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	buf:= make([]byte, 512)
	length, err := conn.Read(buf)
	if err != nil{
		fmt.Println("Error in reading", err.Error())
		return
	}
	fmt.Printf("Received data: %v\n", string(buf[:length]))
	write_len, e :=conn.Write([]byte("nihaoya"))
	if e != nil{
		fmt.Println("Error in reading", e.Error())
	}
	fmt.Printf("Send data: %d\n", write_len)
}
func main() {
	MakeServer("localhost", "12000")
}