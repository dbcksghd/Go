package main

import "fmt"

func main() {
	a := 2
	b := a
	a = 10
	fmt.Println(a, b)
	c := 3
	d := &c
	fmt.Println(c, *d)
	*d = 5
	fmt.Println(c, *d)
}
