package main

import (
	"errors"
	"fmt"
	"net/http"
)

type request struct {
	url    string
	status string
}

var errRequestFailed = errors.New("불러오기 실패")

func main() {

	var results = make(map[string]string)
	urls := []string{
		"https://www.naver.com/",
		"https://www.youtube.com/",
		"https://github.com/",
		"https://www.acmicpc.net/",
		"https://www.google.co.kr/",
	}

	for _, url := range urls {
		fmt.Println(url + "확인 중...")
		result := "성공"
		err := hitURL(url)
		if err != nil {
			result = "실패"
		}
		results[url] = result
	}
	for url, res := range results {
		fmt.Println(url, res)
	}
}

func hitURL(url string) error {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		fmt.Println(err, res.StatusCode)
		return errRequestFailed
	}
	return nil
}
