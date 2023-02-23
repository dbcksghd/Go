package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := &router{make(map[string]map[string]http.HandlerFunc)}

	r.HandleFunc("GET", "/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "welcome!")
	})

	r.HandleFunc("GET", "/what", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "what")
	})

	r.HandleFunc("GET", "/users/:id", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "사용자 검색")
	})

	r.HandleFunc("POST", "/users", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "사용자 생성")
	})
	r.HandleFunc("POST", "/users/:user_id/addresses", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "사용자의 주소 생성")
	})

	http.ListenAndServe(":8080", r)
}
