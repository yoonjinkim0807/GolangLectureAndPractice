package main

import (
	"bufio"
	"log" // Fatal function
	"os"
	// TrimSpace
	// ParseInt
)

func main() {

	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')

	if err != nil { // 에러가 발생하면
		log.Fatal(err)
	}
}
