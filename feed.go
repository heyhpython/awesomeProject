package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const dataFile = "data.json"

type Feed struct {
	Name string `json:"site"`
	Url string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ( []*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err

}

func main() {
	feeds, err := RetrieveFeeds()
	if err!= nil{
		fmt.Println(err)
	}
	for _,feed := range feeds{
		fmt.Printf("name:%s link:%s type:%s", feed.Name, feed.Url, feed.Type )
	}
}
