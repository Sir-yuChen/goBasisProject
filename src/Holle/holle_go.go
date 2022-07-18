package main // 声明 main 包，表明当前是一个可执行程序 注意:每个 Go 应用程序都包含一个名为 main 的包

import "fmt" // 导入内置 fmt 包

func main() { //这里面go语言的语法，定义函数的时候，‘{’ 必须和函数名在同一行，不能另起一行 ,语法严格 否则编译无法通过
	// main函数，是程序执行的入口
	fmt.Println("Hello Golang!") // 在终端打印 Hello Golang!
	fmt.Print("hello, world\n")  //Println  和‘\n’ 作用相同
}
