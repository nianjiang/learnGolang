package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"mycompany.com/test/learngolang/pkg/morestrings"
	// mytype "./myType"
)

func main() {
	printDate()

	fmt.Println("uuid: ", uuid.New().String())

	fmt.Println("fact(3): ", fact(3))

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}

func printDate() {
	fmt.Printf("Hello world!!!\n\t\t%s\n", time.Now())
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
