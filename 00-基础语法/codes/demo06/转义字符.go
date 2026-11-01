package main

import "fmt"

func main() {
	// 练习转义字符
	///n 换行
	fmt.Println("aaa\nbbb")
	fmt.Println("--------------------------")
	//\b 退格
	fmt.Println("aaa\bbbb")
	fmt.Println("--------------------------")
	//\r 光标回到本行的开头，后续输入就会替换掉原本的字符
	fmt.Println("aaa\rbbb")
	fmt.Println("--------------------------")
	// \t制表符
	fmt.Println("aaa\tbbb")
	fmt.Println("--------------------------")
	// \"
	fmt.Println("aaa\"bbb")

}
