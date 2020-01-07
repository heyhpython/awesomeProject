package main

type Options struct {
	para1 int
	para2 string
}

func main() {
	o := Options{
		para1: 1,
		para2: "qweq"}
	PrintParas(o)
}

func PrintParas(o Options) {
	println(o.para1)
	println(o.para2)
}
