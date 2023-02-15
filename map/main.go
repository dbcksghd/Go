package main

import "fmt"

func main() {
	map1 := map[string]string{"이름": "유찬홍", "나이": "18살"}
	map1["생년월일"] = "12월 6일"
	map2 := make(map[int]int)
	fmt.Println(map1)
	fmt.Println(map2)
	value, isTrue := map1["키"]
	map1["생년월일"] = "2월 30일"
	fmt.Println(value, isTrue)
	delete(map1, "나이")
	fmt.Println(map1)
	for key, value := range map1 {
		fmt.Println(key, value)
	}
}
