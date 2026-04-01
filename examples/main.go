package main

import (
	"fmt"
	"time"

	// mytype "./myType"

	"learngolang/examples/basic"
	"learngolang/pkg/mytype"
)

func main() {
	// printDate()

	// fmt.Println("uuid: ", uuid.New().String())

	// goobject.Test()

	// testRecangle()
	testBasic()

}

func testBasic() {
	// basic.PrintHello()
	basic.PrintPointer()
}

func testRecangle() {
	fmt.Println("This is testRecangle......")

	r := mytype.Rectangle{Width: 3, Height: 4}
	fmt.Println(r)

	// mytype.Test()
}

func printDate() {
	fmt.Printf("Hello world!!!\n\t\t%s\n", time.Now())
}
