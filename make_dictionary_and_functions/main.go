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

	addErr := dictionary.Add("키", "165")
	if addErr != nil {
		fmt.Println(addErr)
	}
	searchDef1, searchErr1 := dictionary.Search("키")
	if searchErr1 != nil {
		fmt.Println(searchErr1)
	} else {
		fmt.Println(searchDef1)
	}

	updateErr := dictionary.Update("키", "180")
	if updateErr != nil {
		fmt.Println(updateErr)
	} else {
		searchDef2, _ := dictionary.Search("키")
		fmt.Println(searchDef2)
	}
	dictionary.Delete("키")
	searchDef2, searchErr2 := dictionary.Search("키")
	if searchErr2 != nil {
		fmt.Println(searchErr2)
	} else {
		fmt.Println(searchDef2)
	}
}
