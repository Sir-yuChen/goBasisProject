package main

import "fmt"

/*
	defer语句被用于预定对一个函数的调用。可以把这类被defer语句调用的函数称为延迟函数。
	defer作用：
		释放占用的资源
		捕捉处理异常
		输出日志
	结果
		如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。
*/

// 函数 变量 常量 首字母大小写，go中代表作用域  小写只能在本包中使用  大写可在除本包以外的其它包使用
func dome() {
	fmt.Println("执行了dome函数")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("4")
	defer fmt.Println("3")
}

func main() {
	fmt.Println("执行main")
	dome()
	/*
		执行结果：
			执行main
			执行了dome函数
			4
			3
			2
			1
		从执行结果来看，defer是在所有的正常代码执行完之后执行，如果有多个，这从最后一个执行，【压栈后出栈】
	*/

}
