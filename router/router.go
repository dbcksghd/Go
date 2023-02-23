package main

import (
	"net/http"
	"strings"
)

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
	//http 메소드에 맞는 핸들러를 반복해서 찾는 과정
	for pattern, handler := range r.handlers[request.Method] {
		if ok, _ := match(pattern, request.URL.Path); ok {
			handler(writer, request)
			return
		}
	}
	//못찾으면 에러처리
	http.NotFound(writer, request)
	return
}

func match(pattern, path string) (bool, map[string]string) {
	//둘 다 일치하면 바로 true
	if path == pattern {
		return true, nil
	}

	// "/" 단위로 하나씩 분리
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	// 패턴과 패스를 "/"로 구분한 부분 문자열 집합의 갯수가 다르면 false
	if len(paths) != len(patterns) {
		return false, nil
	}

	//패턴에 맞는 url 파라미터를 담기 위한 맵
	params := make(map[string]string)

	// "/"로 구분된 패턴과 패스의 문자열을 하나씩 비교
	for i := 0; i < len(patterns); i++ {
		switch {
		//case pattern[i] == paths[i]:
		//패턴과 패스의 부분 문자열이 일치하는 경우
		// 바로 다음 루프 수행
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			//패턴이 ':' 문자로 시작하는 경우
			// params에 url params를 담은 후 다음 루프 수행
			params[patterns[i][1:]] = paths[i]
		default:
			// 일치하는 경우가 없으면 false
			return false, nil

		}
	}
	return true, params
}
