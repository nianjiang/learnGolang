package basic

import (
	"fmt"
	"unsafe"
)

func PrintHello() {
	println("Hello, World! From Basic package.")
}

func PrintPointer() {
	var i1 interface{} = 10
	var i2 interface{} = []int{1, 2, 3}

	printInterface(i1)
	printInterface(i2)

}

func printInterface(i interface{}) {
	println("unsafe.Sizeof.", unsafe.Sizeof(i))
	fmt.Printf("Dynamic type: %T, Dynamic value: %v\n", i, i)
}
