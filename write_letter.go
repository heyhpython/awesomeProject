package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bufferReader := bufio.NewReader(os.Stdin)
	var lineCount, wordCount, charCount int
	for {
		line, err := bufferReader.ReadString('\n')
		if err != nil {
			println("something wrong program exit")
			fmt.Printf("lines:%d; words:%d; chars:%d", lineCount, wordCount, charCount)
			break
		}
		if line == "S\n" {
			println("write finish, exit program")
			fmt.Printf("lines:%d; words:%d; chars:%d", lineCount, wordCount, charCount)
			break
		}
		lineCount += 1
		line = strings.TrimSpace(line)
		//line = strings.Replace(line, '\\r', '\ ', -1)
		charCount += len(line)
		//for index, w := range strings.Split(' ')
		//wordCount += len(strings.Split(line, ''))
	}
}
