package main

import (
	"fmt"
	"make_dictionary_and_functions/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"name": "yoochanhong"}
	searchDef, searchErr := dictionary.Search("what")
	if searchErr != nil {
		fmt.Println(searchErr)
	} else {
		fmt.Println(searchDef)
	}

	addErr := dictionary.Add("키", "180")
	if addErr != nil {
		fmt.Println(addErr)
	}
	searchDef1, searchErr1 := dictionary.Search("키")
	if searchErr1 != nil {
		fmt.Println(searchErr1)
	} else {
		fmt.Println(searchDef1)
	}
}
