package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
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

	e := echo.New()

	e.GET("/login", func(c echo.Context) error {
		url := githubConfig.AuthCodeURL("state")
		return c.Redirect(http.StatusTemporaryRedirect, url)
	})

	e.GET("/callback", func(c echo.Context) error {
		code := c.QueryParam("code")
		accessToken, err := githubConfig.Exchange(context.Background(), code)
		if err != nil {
			panic(err)
		}
		return c.JSON(200, accessToken)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
