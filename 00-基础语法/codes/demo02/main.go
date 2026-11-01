package main

import "fmt"

// 下面是全局变量，定义在函数之外的变量
var n8 = 9.8

// 全局变量可以一次性声明
var (
	name   = "han"
	age    = 23
	height = 1.84
)

func main() {
	//定义在{}之内的叫做局部变量

	// 变量的四种声明方式
	fmt.Println("----------------------------------")
	//第一种 ： 变量的使用形式
	var num int = 16
	fmt.Println(num)

	//第二种 ： 指定变量的类型，但是不赋值，使用默认值:0
	var num2 int
	fmt.Println(num2)

	//第三种 ： 不写变量类型 ，根据=后面的值自动判断类型
	var num3 = "tom"
	fmt.Println(num3)

	//第四种： 省掉var ，注意:=不能写成=
	sex := "男"
	fmt.Println(sex)

	fmt.Println("----------------------")
	// 声明多个变量
	var n1, n2, n3, n4 int
	fmt.Println(n1, n2, n3, n4)

	var name, age = "han", 23
	fmt.Println(name)
	fmt.Println(age)

	n5, height := 6.8, 183.2
	fmt.Println(n5, height)
}
