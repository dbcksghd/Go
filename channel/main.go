package main

import "fmt"

func what(c chan string) {
	c <- "what"
}

func main() {
	c := make(chan string)
	go what(c)
	fmt.Println(<-c)

	b := make(chan int, 1)
	func() {
		b <- 5
	}()
	fmt.Println(<-b)

	a := make(chan string, 2)
	a <- "wwwwwhat"
	a <- "wwwhat"
	fmt.Println(<-a)
	fmt.Println(<-a)
	//fmt.Println(<-a)
	fmt.Println(a)
}
