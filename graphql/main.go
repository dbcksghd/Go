package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GraphqlRequest struct {
	Query string `json:"query"`
}

type GraphqlResponse struct {
	Data struct {
		User struct {
			Name      string `json:"name"`
			Following struct {
				Nodes []struct {
					Login   string `json:"login"`
					HtmlUrl string `json:"html_url"`
				} `json:"nodes"`
			} `json:"following"`
			Followers struct {
				Nodes []struct {
					Login   string `json:"login"`
					HtmlUrl string `json:"html_url"`
				} `json:"nodes"`
			} `json:"followers"`
		} `json:"user"`
	} `json:"data"`
}

func main() {
	token := os.Getenv("TOKEN")

	url := "https://api.github.com/graphql"

	query := `query {
		user(login: "yoochanhong") {
				following(first: 100) {
					nodes {
					login
					name
					avatarUrl
					}
				}
				followers(first: 100) {
					nodes {
					login
					name
					avatarUrl
					}
				}
			}
		}
	`

	requestBody := &GraphqlRequest{
		Query: query,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var graphqlResponse GraphqlResponse
	err = json.NewDecoder(resp.Body).Decode(&graphqlResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Following:")
	for _, user := range graphqlResponse.Data.User.Following.Nodes {
		fmt.Println(user.Login)
	}
	fmt.Println("Followers:")
	for _, user := range graphqlResponse.Data.User.Followers.Nodes {
		fmt.Println(user.Login)
	}
}
