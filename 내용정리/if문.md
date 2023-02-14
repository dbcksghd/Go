### if문

##### if문 형식

if문은 다른 언어와 큰 차이는 없다. 

```go
a := 1
if a == 10 {
    fmt.Println("true")
} else {
    fmt.Println("false")
}
```

조건식에 괄호를 넣지 않아도 된다는 점..? 그리고 중괄호를 꼭 넣어주어야 하는 것 같다.

```go
a := 1
if a == 10 {
    fmt.Println("true")
} else fmt.Println("false")
```

else 구문을 이렇게 처리하면 에러가 난다.



##### variable experssion

신기하게도 조건문에서 변수를 만들어낼 수 있다고 한다.

```go
a := 1
if what := a + 10; what == 11 {
    fmt.Println("true")
}
```

if문 안에서 what이라는 변수를 만들고, 그 변수에 대해 조건을 세워서 분기처리하는것도 가능하다.

물론 여기서 생성된 변수는 if문 밖에서는 사용할 수 없다.