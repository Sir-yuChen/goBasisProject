package main

import "fmt"

//Go 值传递和引用传递【传递指针】

/*
	值传递：
		值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
		默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
		总结：实际是值的复制过程，不会影响到源值
*/
//定义相互交换的值
func value(x, y int) int {
	var temp int
	temp = x // 临时变量保存x值
	x = y    //x 赋值为y
	y = temp //x值也就是临时变量中的值，赋值为y 完成值交换
	return temp
}

/*
	引用传递：
		Go 语言中指针是很容易学习的，Go 语言中使用指针可以更简单的执行一些任务
		我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
		Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
		引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
*/
/* 定义交换值函数*/
func valuePointer(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
func main() {

	//使用值传递调用value函数
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("值传递 交换前 a 的值为 : %d\n", a)
	fmt.Printf("值传递 交换前 b 的值为 : %d\n", b)

	/* 通过调用函数来交换值 */
	value(a, b)

	fmt.Printf("值传递 交换后 【源值】a 的值 : %d\n", a)
	fmt.Printf("值传递 交换后 【源值】b 的值 : %d\n", b)
	fmt.Println("==================================")

	//引用传递
	v, h := 10, 20
	//fmt.Printf("v变量的内存地址: %x\n", &v)

	fmt.Printf("引用传递 交换前，v 的值 : %d\n", v)
	fmt.Printf("引用传递 交换前，h 的值 : %d\n", h)

	/*
	 * &v 指向 v 指针，v 变量的地址
	 * &h 指向 h 指针，h 变量的地址
	 */
	valuePointer(&v, &h)

	fmt.Printf("引用传递 交换后，【源值】v 的值 : %d\n", v)
	fmt.Printf("引用传递 交换后，【源值】h 的值 : %d\n", h)
}
