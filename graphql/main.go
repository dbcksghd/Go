package main

import "fmt"

type GithubUser struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
}

func main() {
	fmt.Println("what")
}
