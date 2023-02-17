package main

import (
	"fmt"
	"make_dictionary_and_functions/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"name": "yoochanhong"}
	definition, err := dictionary.Search("what")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
