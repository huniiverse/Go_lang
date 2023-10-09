package main

// 문자열을 숫자로 변환
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" // ParseFloat를 사용하기 위해 "strconv" 패키지를 가져옵니다.
	"strings" // TrimSpace를 사용하기 위해 "strings" 패키지를 가져옵니다.
)

func main() {
	fmt.Print("Enter a grede : ")         // <- 사용자에게 백분율 성적 입력 프롬프트를 출력
	reader := bufio.NewReader(os.Stdin)   // <- 키보드 입력을 읽기 위한 bufio.NewReader를 생성
	input, err := reader.ReadString('\n') // <- 사용자가 엔터키를 누를때까지 입력한 내용을 읽음
	if err != nil {
		log.Fatal(err)
	} // 에러가 발생하면 에러 메시지를 출력하고 종료

	input = strings.TrimSpace(input)            // <- 입력값에서 줄 바꿈 문자를 제거
	grade, err := strconv.ParseFloat(input, 64) // <- 입력 문자열을 float64 값으로 변환
	if err != nil {
		log.Fatal(err)
	} // 위와 동일

	var status string // <- "status"변수는 함수 스코프에서 사용할 수 있도록 이곳에 선언
	if grade >= 60 {
		status = "passing"
	} else {
		status = "falling"
	} // 성적이 60이상이면 "pass", 그 외에는 "fall"으로 설정
	fmt.Println("A grade of", grade, "is", status)
}

/*
import (
	"fmt"
	"bufio"
	"log"
	"os"
)
// error 처리하기
func main() {
	fmt.Print("Enter a grade : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
*/
/*
import (
	"bufio"
	"fmt"
	"os"
)

// 빈 식별자를 사용해 에러 반환 값 무시하기
// 할당은 하지만 사용하지 않는 값에는 빈 식별자(_) 사용
func main() {
	fmt.Println("Enter a grade")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // 에러 값 위치에 빈 식별자 사용
	fmt.Println(input)
}*/
