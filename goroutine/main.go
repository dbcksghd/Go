package main

import (
	"fmt"
	"time"
)

func what() {
	for i := 0; i < 10; i++ {
		fmt.Println("what", i)
		time.Sleep(time.Second)
	}
}

func main() {
	what()
	go what()
	go what()
	time.Sleep(time.Second * 5)
}
