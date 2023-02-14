package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5}
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}

	arr1 = append(arr1, 8)

	arr2 := []int{}
	var arr3 [1][2][3]int
	arr3[0][0][0] = 1

	fmt.Println(len(arr), arr1, arr2, arr3[0][0][0])

	arr4 := []int{0, 1, 2, 3, 4, 5}
	arr5 := []int{}
	arr5 = arr4[1:3]
	arr4 = arr4[:4]
	fmt.Println(arr4, arr5)

	arr6 := []int{1, 2, 3, 4, 5}
	arr7 := make([]int, len(arr6))
	copy(arr7, arr6)
	fmt.Println(arr7)
}
