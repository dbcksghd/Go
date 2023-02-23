package main

import "net/http"

type router struct {
	//키 : http 메소드
	//값 : url별로 실행할 핸들러 함수
	handlers map[string]map[string]http.HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {
	// http 메소드로 등록된 키가 있는지 확인
	m, ok := r.handlers[method]
	if !ok {
		// 만약 없다면 새로운 맵을 만들어서 함수를 키값에 저장
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
	}
	//키로 등록된 맵에 url 패턴과 핸들러 함수 등록
	m[pattern] = h
}

// handlers 맵에서 request.Method와 request.URL.Path에 맞는 핸들러를 실행시키는 함수
func (r *router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if m, ok := r.handlers[request.Method]; ok {
		if h, ok := m[request.URL.Path]; ok {
			//요청한 url에 맞는 핸들러 수행
			h(writer, request)
			return
		}
	}
	http.NotFound(writer, request)
}
