package main //每个文件所在的包，每个Go文件都必须有归属的包

import "fmt" // 引入程序需要的包，为了使用包下的函数，比如PrintLn

func main() {
	fmt.Println("Hello World")
}
