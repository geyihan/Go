package main

import "fmt"

func main() {
	//变量的声明与赋值
	var age int = 16
	//变量的使用
	fmt.Println("age=", age)

	//变量是不可以重复定义的
	//# command-line-arguments
	//./main.go:11:6: age redeclared in this block
	//./main.go:6:6: other declaration of age
	//var age = 18
	//fmt.Println("age=", age)

	//不可以在赋值的时候给予不匹配的类型
	//# command-line-arguments
	//./main.go:18:11: undefined: init
	//var num1 init = 1.34
	var num1 = 1.34
	fmt.Println("num1=", num1)

}
