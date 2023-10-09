package main

import (
	"fmt"
	"strings"
	"time" // time.Time 타입을 사용하기 위해 time 패키지 가져옴
)

func main() {
	var now time.Time = time.Now() // time.Now는 현재 날짜와 시간을 나타내는 time.Time 값 반환
	var year int = now.Year()      // time.Time 값은 연도를 반환하는 Year 매서드를 가지고 있음
	fmt.Println(year)
	//------------------------------------------------------------------------
	broken := "G# r#cks"
	replace := strings.NewReplacer("#", "o") // '#'을 'o'으로 치환하도록 설정된 string.Replacer를 반환
	fixed := replace.Replace(broken)         // string.Replacer의 Replace 매서드를 호출해 치환할 문자열 전달
	fmt.Println(fixed)                       // Replace 매서드에서 반환된 문자열 출력
}

// 메소드 호출하기
// now.Year() -> 값.매서드명
// replacer.Replace(broken) -> 값.매서드명
// .(dot, 점)은 오른쪽에 있는 무언가가 왼쪽의 무언가에 속해 있음을 나타냄
