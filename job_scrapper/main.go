package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

var baseURL string = "https://browse.auction.co.kr/search?keyword=%EB%A7%A5%EB%B6%81&k=31"

func main() {
	totalPages := getPages()
	for i := 1; i <= totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&p=" + strconv.Itoa(page)
	fmt.Println(pageURL)
}

func getPages() int {
	pages := 0
	response, errorMessage := http.Get(baseURL)
	checkErr(errorMessage)
	checkStatusCode(response)
	document, err := goquery.NewDocumentFromReader(response.Body)
	checkErr(err)
	document.Find(".component--pagination").Each(func(i int, selection *goquery.Selection) {
		pages = selection.Find("a").Length()
	})
	defer response.Body.Close()
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
