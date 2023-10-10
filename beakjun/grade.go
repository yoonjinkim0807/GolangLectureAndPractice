package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("점수를 입력하세요")
	gradeRd := bufio.NewReader(os.Stdin)
	gradIn, err := gradeRd.ReadString('\n')

	var a

	fmt.Println(a)

	if err != nil {
		log.Fatal(err)
	}
	gradIn = strings.TrimSpace(gradIn)
	gradInConv, err := strconv.Atoi(gradIn)
	if err != nil {
		log.Fatal(err)
	} else {
		if gradInConv >= 90 {
			fmt.Println("A")
		} else if gradInConv >= 80 {
			fmt.Println("B")
		} else if gradInConv >= 70 {
			fmt.Println("C")
		} else if gradInConv >= 60 {
			fmt.Println("D")
		} else {
			fmt.Println("F")
		}
	}

}
