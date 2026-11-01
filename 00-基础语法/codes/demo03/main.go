package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 = 230
	fmt.Println(num1)

	var num2 uint8 = 255
	fmt.Println(num2)

	var num3 = 28
	fmt.Printf("num3的类型是： %T", num3)
	fmt.Println(unsafe.Sizeof(num3))
}
