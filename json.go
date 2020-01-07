package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func dump() {
	a1 := &Address{"private", "上海", "中国"}
	a2 := &Address{"work", "武汉", "中国"}

	card := VCard{"yuhao", "He", []*Address{a1, a2}, "none"}

	js, _ := json.Marshal(card)
	fmt.Printf("Json format: %s", js)
	file, _ := os.OpenFile("data.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(card)
	if err != nil {
		log.Println("error in encoding")
	}
}

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func load() {
	//var m FamilyMember
	//b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	//err := json.Unmarshal(b,&m)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(m)
	var card VCard
	f, _ := os.Open("data.json")
	decoder := json.NewDecoder(f)
	err := decoder.Decode(&card)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(card)
	fmt.Println(card.Addresses[0])
}

func main() {
	dump()
	load()
}
