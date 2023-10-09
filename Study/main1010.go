package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
}

// 메소드 호출하기
//
