package main

import (
	"fmt"
	"log"
	"net/http"
)

var baseURL string = "https://www.google.com/search?q=flutter&rlz=1C5CHFA_enKR1035KR1035&oq=flutter&aqs=chrome..69i57j69i59l3j69i60l4.3615j0j4&sourceid=chrome&ie=UTF-8"

func main() {
	pages := getPages()
	fmt.Println(pages)
}

func getPages() int {
	response, errorMessage := http.Get(baseURL)
	checkErr(errorMessage)
	checkStatusCode(response)
	return 0
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
