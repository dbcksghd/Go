package main

import (
	"echo/utils"
	"encoding/json"
	"io/ioutil"
	"os"
)

type m struct {

}

func getAuth() string {
	data, err := os.Open("/Users/yoochanhong/Desktop/github/Go/echo/database/asdf.json")
	utils.CheckErr(err)

	var auth
	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, auth)
	return auth
}
