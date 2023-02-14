package main

import "fmt"

func main() {
	var a int = 1
	switch {
	case a == 1:
		fmt.Println(1)
		fallthrough
	case a == 2:
		fmt.Println(2)
	case a == 3:
		fmt.Println(3)
	}

	switch v := 100; {
	case v > 50:
		fmt.Println("50 이상")
	case v > 10:
		fmt.Println("10 이상")
	case v > 0:
		fmt.Println("0 이상")

	}
}
