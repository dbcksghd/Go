package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("불러오기 실패")

func main() {

	var results = make(map[string]string)
	ch := make(chan requestResult)
	urls := []string{
		"https://www.naver.com/",
		"https://www.youtube.com/",
		"https://github.com/",
		"https://www.acmicpc.net/",
		"https://www.google.co.kr/",
	}

	for _, url := range urls {
		go hitURL(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		result := <-ch
		results[result.url] = result.status

	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, ch chan<- requestResult) {
	res, err := http.Get(url)
	status := "성공"
	if err != nil || res.StatusCode >= 400 {
		status = "실패"
	}
	ch <- requestResult{url: url, status: status}
}
