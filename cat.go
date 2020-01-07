package main

import (
	"flag"
	"fmt"
	"os"
)

func cat(r *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, _ := r.Read(buf[:]); {
		//读出类的内容都在buf里  返回的是读到的字节数
		case nr < 0:
			fmt.Println("cat error read", err.Error())
		case nr == 0:
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Println("cat error", ew.Error())
			}

		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(os.Stdin)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		defer f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s", err, flag.Arg(i))
			continue
		}
		cat(f)

	}

}
