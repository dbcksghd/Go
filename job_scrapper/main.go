package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

var baseURL string = "https://github.com/search?q=what&type=users"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)
}

func getPages() int {
	pages := 0
	response, errorMessage := http.Get(baseURL)
	checkErr(errorMessage)
	checkStatusCode(response)
	document, errorM := goquery.NewDocumentFromReader(response.Body)
	checkErr(errorM)
	document.Find(".codesearch-pagination-container").Each(func(i int, selection *goquery.Selection) {
		pages = selection.Find("a").Length()
	})
	defer response.Body.Close() // 함수가 끝났을 때 실행할꺼
	return pages
}

func checkErr(errorMessage error) {
	if errorMessage != nil {
		log.Fatalln(errorMessage) // 만약 에러가 있다면 로그찍고 프로그램 끝내버리기
	}
}

func checkStatusCode(response *http.Response) {
	if response.StatusCode != 200 {
		log.Fatalln("에러코드 : ", response.StatusCode) // 200이 아니여도 프로그램 끝내기
	}
}
