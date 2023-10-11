package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n) // 리턴되는 값을 지정함
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil //다중 리턴이기때문에 2이상의 수가 왔으니 정상값인 nil을반환
		}

	}

	return true, nil // 소수 판정 값, 정상데이터(에러)
}

func prime(number int) {
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

func primeRange(a int, b int) {
	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if p {
			fmt.Print(i, " ")
		}
	}
}

// 소수 판정 및 구간 소수 판정 프로그램 v1.7
func main() {
	var menu int

	for true {
		fmt.Print("1) 소수판정 2) 구간 소수판정 : ")
		_, err := fmt.Scanln(&menu) // n은 입력받은 개수, n을 _로 하면 사용을 하지않음(return X)

		if err != nil {
			log.Fatal(err) // int로 number를 정의 했기 때문에 number에 int를 제외한 값이 들어가면 오류
		}

		switch menu {
		case 1:
			var in int
			fmt.Print("정수 1개입력:")
			_, err := fmt.Scanln(&in) // n은 입력받은 개수, n을 _로 하면 사용을 하지않음(return X)

			if err != nil {
				log.Fatal(err) // int로 number를 정의 했기 때문에 number에 int를 제외한 값이 들어가면 오류
			}
			prime(in)
		case 2:
			var n1, n2 int
			fmt.Print("정수 2개입력:")
			_, err := fmt.Scanln(&n1, &n2) // n은 입력받은 개수, n을 _로 하면 사용을 하지않음(return X)

			if err != nil {
				log.Fatal(err) // int로 number를 정의 했기 때문에 number에 int를 제외한 값이 들어가면 오류
			}
			primeRange(n1, n2)
		default:
			fmt.Print("프로그램을 종료합니다.")
			os.Exit(1)
		}

	}

}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {
// 	if n < 2 {
// 		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n) // 리턴되는 값을 지정함
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return false, nil //다중 리턴이기때문에 2이상의 수가 왔으니 정상값인 nil을반환
// 		}

// 	}

// 	return true, nil // 소수 판정 값, 정상데이터(에러)
// }

// // 구간 소수 판정 프로그램 v1.4 inPrime 함수 안의 변수(prime)를 줄이고 return 구문추가, break 제거
// func main() {
// 	var a, b int
// 	//2,20
// 	// 2 3 5 7 11 13 ... 19

// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&a, &b) // n은 입력받은 개수, n을 _로 하면 사용을 하지않음(return X)

// 	if err != nil {
// 		log.Fatal(err) // int로 number를 정의 했기 때문에 number에 int를 제외한 값이 들어가면 오류
// 	}

// 	if a > b {
// 		temp := a
// 		a = b
// 		b = temp
// 	}

// 	for i := a; i <= b; i++ {
// 		p, err := isPrime(i)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(0)
// 		}
// 		if p {
// 			fmt.Print(i, " ")
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {
// 	prime := true

// 	if n < 2 {
// 		return false, fmt.Errorf("%d는(은) 소수가 아닙니다.", n) // 리턴되는 값을 지정함
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			prime = false
// 			break
// 		}

// 	}

// 	return prime, nil // 소수 판정 값, 정상데이터(에러)
// }

// // 구간 소수 판정 프로그램 v1.3
// func main() {
// 	var a, b int
// 	//2,20
// 	// 2 3 5 7 11 13 ... 19

// 	fmt.Print("정수 입력 : ")
// 	_, err := fmt.Scanln(&a, &b) // n은 입력받은 개수, n을 _로 하면 사용을 하지않음(return X)

// 	if err != nil {
// 		log.Fatal(err) // int로 number를 정의 했기 때문에 number에 int를 제외한 값이 들어가면 오류
// 	}

// 	if a > b {
// 		temp := a
// 		a = b
// 		b = temp
// 	}

// 	for i := a; i <= b; i++ {
// 		p, err := isPrime(i)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(0)
// 		}
// 		if p {
// 			fmt.Print(i, " ")
// 		}
// 	}
// }
