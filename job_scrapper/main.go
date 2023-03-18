package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type extractedJob struct {
	id       string
	location string
	title    string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {

	var jobs []extractedJob
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob

	pageURL := baseURL + "&start" + strconv.Itoa(page*50)
	response, errorMessage := http.Get(pageURL)
	checkErr(errorMessage)
	checkStatusCode(response)
	document, err := goquery.NewDocumentFromReader(response.Body)
	checkErr(err)
	defer response.Body.Close()

	searchCards := document.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, selection *goquery.Selection) {
		job := extractJob(selection)
		jobs = append(jobs, job)
	})
	return jobs
}

func getPages() int {
	pages := 0
	response, errorMessage := http.Get(baseURL)
	checkErr(errorMessage)
	checkStatusCode(response)
	document, err := goquery.NewDocumentFromReader(response.Body)
	checkErr(err)
	document.Find(".pagination").Each(func(i int, selection *goquery.Selection) {
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

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), "")
}

func extractJob(selection *goquery.Selection) extractedJob {
	id, _ := selection.Attr("data-jk")
	title := cleanString(selection.Find(".title>a").Text())
	location := cleanString(selection.Find(".sjcl").Text())
	salary := cleanString(selection.Find(".salatyText").Text())
	summary := cleanString(selection.Find(".summary").Text())
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}
