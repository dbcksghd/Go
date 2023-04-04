package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

var baseurl = "https://api.github.com/users/"

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "what")
	})

	//e.GET("/unfollowing", func(c echo.Context) error {
	//	m := make(map[string]int)
	//	followerCh := make(chan []User)
	//	followingCh := make(chan []User)
	//
	//	userName := c.QueryParam("userName")
	//
	//	getFollowingUserList(userName, followingCh)
	//	getFollowerUserList(userName, followerCh)
	//	var list []User
	//	list <- followingCh
	//	for user := range followingCh {
	//		m[user] = 1
	//	}
	//	for user := range followerCh {
	//		if m[user.Login] != 1 {
	//			list = append(list, user)
	//		}
	//	}
	//	return c.JSON(200, list)
	//})

	e.GET("/unfollower", func(c echo.Context) error {
		m := make(map[string]int)
		followerCh := make(chan []User)
		followingCh := make(chan []User)
		//userSet1Ch := make(chan User)
		userName := c.QueryParam("userName")
		getFollowUserList(userName, "folowing", followingCh)
		getFollowUserList(userName, "followers", followerCh)

		list := <-followingCh
		for i, user := range list {
			go userSet1(user, m, i)
		}
		//for user := range followingCh {
		//	if m[user.Login] != 1 {
		//		list = append(list, user)
		//	}
		//}
		//return c.JSON(200, list)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func userSet1(user User, m map[string]int, i int) {
	m[user.Login] = 1
	fmt.Println(i)
}

// 내가 팔로잉한 사람들을 긁어오는 함수
func getFollowUserList(userName string, follow string, ch chan<- []User) {
	for i := 1; ; i++ {
		pageURL := baseurl + userName + "/" + follow + "?per_page=100&page=" + strconv.Itoa(i)
		req, err := http.NewRequest("GET", pageURL, nil)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Authorization", "Bearer"+token)
		client := &http.Client{}

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		var following []User
		err = json.Unmarshal(body, &following)
		if err != nil {
			log.Fatalf("Failed to decode following: %v", err)
		}
		ch <- following

		if len(following) < 100 {
			break
		}
	}
	close(ch)
}
