package main

import (
	"flag"
	"fmt"
	"os"
)

var NLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	NewLine = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()
	var s = ""
	for i := 0; i < flag.NArg(); i++ {
		if i >= 0 {
			s += " "
			if *NLine {
				s += NewLine
			}
			fmt.Println(flag.Arg(i))
			s += flag.Arg(i)
		}
	}
	os.Stdout.WriteString(s)

}
