package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time = time.Now()
	fmt.Println(t.Format(time.ANSIC))
}
