package main

import ( //모듈 가져오기
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	//var a int // declaration
	//a = 7     //assign value

	// var a int = 7 //decaration & assign value

	//var a = 7 // literal을 보고 자료형을 결정 a의 결정된 타입이 변형 되지 않음(파이썬과 다른점)

	a := 7 // := 단축 연산자
	fmt.Println(a, reflect.TypeOf(a))
	// b := 8.34 => 기본적으로 8byte로 할당 됨 float64
	var b float32 = 8.34
	fmt.Println(a * int(b))
	fmt.Println(float32(a) > b)
	//fmt.Println(b, reflect.TypeOf(b))

	var c7 string // 빈문자열 , 변수명은 문자로 시작하고 숫자는 첫글자 뒤에 와야한다.
	var d int     // 정수 0
	var e bool    // 기본값 false
	var f float64 // 실수 0
	var G = 99    // 대문자로 시작하면 외부로 공개 할 수 있다. 따라서 패키지를 불러올 수 있게 된다. 소문자로 시작하면 같은 패키지 안에서만 사용가능

	fmt.Println(c7, d, e, f, G)
	// koreanzombie := "정찬성" 문법적으로 문제 없지만 camel case를 사용하는 커뮤니티관례에 맞지 않는다.

	koreanZombie := "정찬성"
	fmt.Println(koreanZombie)

	fmt.Println('Z', '2', '\n', '김', '인', '하') // 90 50 10 44608 51064 54616 유니코드 숫자로 출력 rune literals(int32)'' 을 사용
	fmt.Println(reflect.TypeOf('2'), reflect.TypeOf(2), reflect.TypeOf("hi"), reflect.TypeOf(4.99), reflect.TypeOf(false))
	fmt.Println(math.Floor(2.17), math.Ceil(2.17), math.Sqrt(16))
	fmt.Println(strings.Title("open source \t programming \n \"go\"!")) // Title 영문자의 첫 글자를 대문자로 출력
	//fmt.Println(strings.Title("오픈소스 프로그래밍"))

	//fmt.Println("OpenSource Programing")
}
