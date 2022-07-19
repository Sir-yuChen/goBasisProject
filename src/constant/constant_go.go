package main

import (
	"fmt"
	"unsafe" //导入多个包
)

/*	常量
	1. 常量是一个简单值的标识符，在程序运行时，不会被修改的量
	2. 常量中的数据类型只可以是【布尔型、数字型（整数型、浮点型和复数）和字符串型】
	常量的定义语法：
		const identifier [type] = value
	[可省略type类型的定义，编译器可以根据变量的值来推断其类型]
*/

func main() {

	//显式类型定义
	const falg bool = true
	fmt.Printf("显式类型定义 falg = %t\n", falg)
	//隐私类型定义
	const falg2 = "HOLLE"
	fmt.Println("隐私类型定义 falg2 = ", falg2)

	//多重赋值
	const a, b, c = 1, 2.8, "Holle 你好"
	fmt.Println("多重赋值：", a, b, c)

	//常量，作为枚举值定义
	const (
		SUCCESS = true  //成功
		FLASE   = false //失败
		UNKNOW  = 0     //未知
		FEMALE  = 1     //女
		MALE    = 2     //男
		S       = "abc"
	)
	fmt.Println("定义枚举值：", SUCCESS, FLASE, UNKNOW, FEMALE, MALE)

	//常量相关方法： len(), cap(), unsafe.Sizeof()常量计算表达式的值 【常量表达式中，函数必须是内置函数，否则编译不过】
	fmt.Println("len()= ", len(S))
	fmt.Println("unsafe.Sizeof()= ", unsafe.Sizeof(S))
	//字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。

}
