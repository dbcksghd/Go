package main

import "fmt"

func main() {
	a := 1
	if a == 10 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if what := a + 10; what == 11 {
		fmt.Println("true")
	}
}
