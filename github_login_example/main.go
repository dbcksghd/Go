package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"os"
)

func main() {
	client_id := os.Getenv("client_id")
	client_secret := os.Getenv("client_secret")

	githubConfig := &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
}
