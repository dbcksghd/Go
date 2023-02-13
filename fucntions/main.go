package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper1(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func print(words ...string) {
	fmt.Println(words)
}

func whatDefer() {
	defer fmt.Println("함수 끝")
	fmt.Println("함수 실행")
}

func main() {
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("유찬홍")
	fmt.Println(totalLength, upperName)
	fmt.Println(lenAndUpper1("이설하"))
	print("1", "2", "3", "4", "5")
	whatDefer()
}
