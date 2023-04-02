package main

import (
	"encoding/json"
	"fmt"
	"json_parsing/utils"
	"net/http"
)

type Photo struct {
	AlbumId      int
	Id           int
	Title        string
	Url          string
	ThumbnailUrl string
}

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/photos")
	utils.CheckErr(err)
	var photos []Photo
	err = json.NewDecoder(response.Body).Decode(&photos)
	utils.CheckErr(err)
	for _, photo := range photos {
		fmt.Println(photo.Title)
	}
}
