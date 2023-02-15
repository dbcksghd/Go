package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

type what struct {
	data map[int]string
}

func makeWhat() *what {
	a := what{}
	a.data = map[int]string{}
	return &a
}

func main() {
	favFood := []string{"김치", "라면"}
	a := person{"유찬홍", 18, favFood}
	var b = person{}
	b = person{
		name:    "유찬홍",
		age:     18,
		favFood: []string{"김치", "라면"},
	}
	c := new(person)
	fmt.Println(a, b, c)
	d := makeWhat()
	d.data[0] = "what"
	fmt.Println(d.data[0])
}
