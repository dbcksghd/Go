package main

import "fmt"

func add(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	var sum int = 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum > 10 {
		sum -= 10
	}
	fmt.Println(sum)
	//for {
	//	fmt.Println("무한루프")
	//}
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
