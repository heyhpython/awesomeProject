package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:12000")
	if err != nil{
		fmt.Println("error in connect", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("input your message:")
	msg, _ := inputReader.ReadString('\n')
	data := make([]byte, 512)
	var length int

	_, err = conn.Write([]byte(msg))
	length, err = conn.Read(data)
	if err != nil{
		fmt.Println(err.Error())
		return
	}else if length != 0{
		fmt.Printf("receive reply:%v", string(data[:length]))
	}


}