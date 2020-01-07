package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func wr_gz(name string) error {
	var r *bufio.Reader
	fi, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			//fmt.Println(err)
			break
		} else if err != nil {
			fmt.Println(err)
		}
		fmt.Println(line)

	}
	return nil
}

func main() {
	err := wr_gz("defer_.go.gz")
	if err != nil {
		println(err)
	}
	return
	inputFile, inputErr := os.Open("copy_.go")
	if inputErr != nil {
		fmt.Println("an error occured", inputErr)
		return
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		s, readErr := inputReader.ReadString('\n')
		fmt.Println(s)
		if readErr == io.EOF {
			return
		}
	}

}
