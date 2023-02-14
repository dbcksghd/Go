### switch case

##### switch case문의 형태

사실 얘도 전체적인건 크게 다를게 없다.

```go
a := 1
switch a {
case 1:
    fmt.Println(1)
case 2:
	fmt.Println(2)
case 3:
	fmt.Println(3)
}
```

switch 문 뒤에 Expression를 설정하고, case별로 하나하나 분기처리를 해준다.

그런데 Go는 편리하게도 case 하나하나마다 마지막에 break를 자동으로 걸어준다.

만약 내가 여러개의 case를 조건으로 가지고 싶다면, fallthrough 키워드를 적어주면 된다.

```go
a := 1
switch a {
case 1:
    fmt.Println(1)
    fallthrough
case 2:
	fmt.Println(2)
case 3:
	fmt.Println(3)
}
```

실행결과 :

> 1 2

case 1에서 끝나야 하지만 2까지 같이 출력된다.



##### Expression 생략

switch 다음의 Expression을 생략하고, case에서 조건을 나눌 수 있다.

위의 코드를 바꿔보면 이렇게 된다.

```go
a := 1
switch {
case a == 1:
    fmt.Println(1)
    fallthrough
case a == 2:
	fmt.Println(2)
case a == 3:
	fmt.Println(3)
}
```

if else로도 나타낼 수 있지만, 이렇게 코드를 적으면 더 깔끔해보이고 좋다.



##### variable experssion

if문과 마찬가지로 조건문에서 변수를 만들어 낼 수 있다.

```go
switch v := 100; {
case v > 50:
	fmt.Println("50 이상")
case v > 10:
	fmt.Println("10 이상")
case v > 0:
	fmt.Println("0 이상")

}
```

실행결과 : 

> 50 이상