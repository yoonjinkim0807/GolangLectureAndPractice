package main

import "fmt"

/*-shadowing 자료타입,패키지,함수를 변수명으로 사용했을 시 문제점
자료타입,패키지,함수가 정상적으로 작동하지 않음*/
func main() {

	// 정상적인 변수명 선언
	var test1 float64 = 9.1
	var test2 float64 = 7.9
	var univ string = "inha"

	var f1 string = "functions"
	var f2 = append([]string{}, "함수")

	fmt.Println(test1 + test2)
	fmt.Println(univ)
	fmt.Println(f1)
	fmt.Println(f2)

	// 자료타입을 변수명으로 사용
	// var float64 float64 = 9.1
	// var test float64 = 7.9
	// fmt.Println(float64)
	// 패키지를 변수명으로 사용
	// var fmt string = "inha"
	// fmt.Println()
	// 함수를 변수명으로 사용
	// var append string = "functions"
	// var f = append([]string{},"함수")
}

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	//"time" //seed 생성용 패키지
// 	//"math/rand"
// 	"bufio"
// 	"log"
// 	"os"
//)

// package main

// import (
// 	"fmt"
// 	"log"
// )

// // 입력(0,1처리)된 수의 소수 판정 프로그램 v0.9
// func main() {
// 	var number int
// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&number) // n은 입력받은 개수, n을 _로 하면 사용을 하지않음(return X)

// 	if err != nil {
// 		log.Fatal(err) // int로 number를 정의 했기 때문에 number에 int를 제외한 값이 들어가면 오류
// 	}

// 	for number < 2{
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 		os.Exit(0) //0은 정상(성공적인 종료)이라는 뜻이며 for문을 멈춰줌 정상적으로 for문 종료
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}

// 	}
// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}

// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	//"time" //seed 생성용 패키지
// 	//"math/rand"
// 	"bufio"
// 	"log"
// 	"os"
// )

// package main

// import (
// 	"fmt"
// 	"log"
// )

// // 입력(fmt 패키지의 Scanln())된 수의 소수 판정 프로그램 v0.8
// func main() {
// 	var number int
// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&number) // n은 입력받은 개수, n을 _로 하면 사용을 하지않음(return X)

// 	if err != nil {
// 		log.Fatal(err) // int로 number를 정의 했기 때문에 number에 int를 제외한 값이 들어가면 오류
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}

// 	}
// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}

// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	//"time" //seed 생성용 패키지
// 	//"math/rand"
// 	"bufio"
// 	"log"
// 	"os"
// )

// // 입력된 수의 소수 판정 프로그램 v0.7
// func main() {

// 	fmt.Print("정수 입력 : ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n') //엔터키 앞까지 입력받음

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	in = strings.TrimSpace(in)      //문자열 앞뒤 공란 제거
// 	number, err := strconv.Atoi(in) // 정수로 변경
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}

// 	}
// 	if isPrime {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}

// }
