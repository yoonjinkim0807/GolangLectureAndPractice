// 사용자 정의 함수
package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	prime := true

	if n < 2 {
		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n) // 리턴되는 값을 지정함
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break
		}

	}

	return prime, nil // 소수 판정 값, 정상데이터(에러)
}

// 소수 판정 프로그램 v1.1 : 함수 적용,error 리턴
func main() {
	var number int

	fmt.Print("정수 입력 : ")
	_, err := fmt.Scanln(&number) // n은 입력받은 개수, n을 _로 하면 사용을 하지않음(return X)

	if err != nil {
		log.Fatal(err) // int로 number를 정의 했기 때문에 number에 int를 제외한 값이 들어가면 오류
	}

	p, err := isPrime(number)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if p {
		fmt.Println(number, " 는(은) 소수입니다.")
	} else {
		fmt.Println(number, " 는(은) 소수가 아닙니다.")
	}
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool) {
// 	prime := true

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			prime = false
// 			break
// 		}

// 	}

// 	return prime //true를 리턴하면 소수, false면 소수아님
// }

// // 소수 판정 프로그램 v1.0 : 함수 적용
// func main() {
// 	var number int

// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&number) // n은 입력받은 개수, n을 _로 하면 사용을 하지않음(return X)

// 	if err != nil {
// 		log.Fatal(err) // int로 number를 정의 했기 때문에 number에 int를 제외한 값이 들어가면 오류
// 	}

// 	if number < 2 {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 		os.Exit(0) //0은 정상(성공적인 종료)이라는 뜻이며 for문을 멈춰줌 정상적으로 for문 종료
// 	}

// 	if isPrime(number) {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// after (multi return)
// package main

// import (
// 	"fmt"
// )

// func processScore(kor int, eng int, math int) (int, int) { // 매개변수 옆에 ()안에 리턴값은 타입을 지정해준다.
// 	totalScore := kor + eng + math
// 	average := totalScore / 3

// 	return totalScore, average
// 	//fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", name, totalScore, average)
// }

// func main() {
// 	var t int
// 	var a int
// 	t, a = processScore(100, 90, 93) // return 된 토탈값을 t 평균값을 a에 반환
// 	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", "홍길동", t, a)
// 	t, a = processScore(89, 91, 92)
// 	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", "홍길순", t, a)
// }

// after
// package main

// import (
// 	"fmt"
// )

// func processScore(name string, kor int, eng int, math int) {
// 	totalScore := kor + eng + math
// 	average := totalScore / 3

// 	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다.\n", name, totalScore, average)
// }

// func main() {
// 	processScore("홍길동", 100, 90, 93)
// 	processScore("홍길동", 89, 91, 92)
// }

//before
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	kor := 100
// 	eng := 90
// 	math := 93
// 	name := "홍길동"

// 	fmt.Println(name, "의 총점은 ", (kor + eng + math), "입니다. 평균은", (kor+eng+math)/3.0)

// 	kor = 99
// 	eng = 91
// 	math = 92
// 	name = "홍길순"

// 	fmt.Println(name, "의 총점은 ", (kor + eng + math), "입니다. 평균은", (kor+eng+math)/3.0)
// 	// ...
// }

// package main

// import (
// 	"fmt"
// )

// func sayHello() {
// 	fmt.Println("Hello~")
// }

// func main() {
// 	sayHello()
// }
