package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("불러오기 실패")

func main() {
	urls := []string{
		"https://www.naver.com/",
		"https://www.youtube.com/",
		"https://github.com/",
		"https://www.acmicpc.net/",
		"https://www.google.co.kr/",
	}
	for _, url := range urls {
		fmt.Println("checking ", url)
		hitURL(url)
	}
}

func hitURL(url string) error {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
