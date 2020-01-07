package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	err := http.ListenAndServe("localhost:13000", nil)
	checkError(err, "")
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	recv := make([]byte, 512)
	length, err := req.Body.Read(recv)
	//checkError(err, req.RequestURI)
	fmt.Println(err.Error())
	fmt.Printf("get msg from client:%s, length:%d", string(recv[:]), length)
	_, err = w.Write([]byte("nihaoya"))
	checkError(err, req.Host)
}

func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}