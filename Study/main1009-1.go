package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("안녕하세요")
// 	var number int
// 	number = 4
// 	fmt.Println(number)
// }

func main() {
	// fmt.Println("안녕하세요")
	// number := 4
	// fmt.Println(number)
	// fmt.Println(reflect.TypeOf(float64(number)))
	var price int = 100
	fmt.Println("price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars")

	var Funds int = 120
	fmt.Println(Funds, "dollars avi")
}

// 변수 선언하기
// var quantity int -> var [변수명] [변숫값의s 타입]
// 변수 선언 -> 변수에 값 할당 -> 변수 사용하기
// 변수에 들어갈 값을 알고 있으면 동시에 할당 가능 -> var number int = 4
// 값을 할당하지 않고 선언한 변수는 타입에 대한 제로값으로 초기화

// 변수 단축 선언 :=
// number := 4
// 단축 변수 선언을 사용하면 변수의 타입은 할당한 값의 타입으로 자동 지정됨

// 네이밍 규칙
// 변수, 함수, 타입의 이름이 대문자면 외부로 노출되어 접근 가능
// 소문자는 외부로 노출되지 않음

// 타입 변환
// 변환하는 타입을 앞에 쓰고 뒤에 변환하려는 값을 괄호로 감싸 주면 됩니다.
// myInt := 2
// float(myInt)

// number := 4
// fmt.Println(number)
// fmt.Println(reflect.TypeOf(float64(number)))
