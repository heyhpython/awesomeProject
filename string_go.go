package main

import "strings"

func main() {
	s := "hello"
	println(len(s))
	println(strings.HasPrefix(s, "he"))
	println(strings.HasSuffix(s, "lo"))
	println(strings.Contains(s, "ll"))
	println(strings.Index(s, "ll"))
	println(strings.LastIndex(s, "l"))
	println(strings.IndexRune(s, 12))
	println(strings.Replace(s, "ll", "00", 1))
	println(strings.Count(s, "l"))
	println(strings.ToUpper(s))

	// 字符串的分割
	for _, val := range strings.Fields(s) {
		println(val)
	}
	for _, val := range strings.Split(s, "l") {
		println(val)
	}

}
