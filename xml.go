package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type address struct {
	Type    string
	City    string
	Country string
}

type vCard struct {
	FirstName string
	LastName  string
	Addresses []*address
	Remark    string
}

func dumps() {
	a1 := &address{"private", "上海", "中国"}
	a2 := &address{"work", "武汉", "中国"}

	card := vCard{"yuhao", "He", []*address{a1, a2}, "none"}
	x, _ := xml.Marshal(card)
	fmt.Println(x)
	file, _ := os.OpenFile("data.xml", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	encoder := xml.NewEncoder(file)
	err := encoder.Encode(card)
	if err != nil{
		fmt.Println(err)
	}
}

func loads() {
	var card vCard
	f, _ := os.Open("data.xml")
	decoder := xml.NewDecoder(f)
	err := decoder.Decode(&card)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(card)
	fmt.Println(card.Addresses[0])
}
func main() {
	//dumps()
	loads()
}
