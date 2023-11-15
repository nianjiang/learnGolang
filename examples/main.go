package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	// mytype "./myType"
)

func main() {
	printDate()

	fmt.Println("uuid: ", uuid.New().String())
}

func testRecangle() {
	// r := mytype.Rectangle{width: 3, height: 4}
	// fmt.Println(r)
}

func printDate() {
	fmt.Printf("Hello world!!!\n\t\t%s\n", time.Now())
}
