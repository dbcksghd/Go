package main

import "fmt"

func main() {
	const name string = "유찬홍"
	fmt.Println(name)
	//name = "이정호" // 상수값은 마음대로 바꿀수 없음
	//fmt.Println(name)
	var name1 string = "조병진"
	fmt.Println(name1)
	name2 := "박준수"
	fmt.Println(name2)
	name3 := false
	//name3 = "김연우" <-- 타입추론으로 만들어진 변수는 다른 자료형으로 바꿀수 없음
	fmt.Println(name3)
}
