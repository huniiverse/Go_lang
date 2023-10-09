package main // package란 유사한 기능을 수행하는 코드들의 모음

import (
	"fmt" // fmt 패키지에 있는 문자열 서식 코드를 사용하기 위해 패키지를 가져오고 있음을 알려줌
	"math"
	"strings"
)

func main() {
	fmt.Println("Hello Go~") // fmt 패키지에 있는 Println이라는 함수를 사용
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
}

//Go file의 기본 형식
// 1. package 절
// 2. import 문
// 3. 실제 코드
