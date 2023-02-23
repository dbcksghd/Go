package main

import "net/http"

type router struct {
	//키 : http 메소드
	//값 : url별로 실행할 핸들러 함수
	handlers map[string]map[string]http.HandlerFunc
}
