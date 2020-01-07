package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func sha1_() {
	hasher := sha1.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("%x\n", hasher.Sum(b))
	fmt.Printf("%d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("we shall overcome")
	n, err := hasher.Write(data)
	if n!=len(data) || err != nil{
		fmt.Println("error %v/%v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("%x\n", checksum)

}
func main() {
	sha1_()
}