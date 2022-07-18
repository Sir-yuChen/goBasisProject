package main

import "fmt"

//GO 声明变量

//声明 多个全局变量
var ( //分解写法
	a bool
	b int
	c = "全局变量"
)

func main() {

	//第一种：使用默认值
	var number int
	fmt.Printf("声明变量  方式一[默认值] number = %d\n", number)

	//第二种方式： 声明并赋值
	var mun int = 10
	fmt.Printf("声明变量  方式二[声明并赋值] mun = %d\n", mun)

	//第三种方式： 省略数据类型赋值 【go语言会自动匹配数据类型】
	var name = "张三" // 注意：字符串不可以用单引号
	fmt.Printf("声明变量  方式三 省略数据类型赋值 【go语言会自动匹配数据类型】 name = %s\n", name)

	//第四种方式： 省略var声明关键字 使用‘：’
	d := 3.14 //注意：冒号要和等号紧挨着
	fmt.Printf("声明变量  方式四 【省略var声明关键字 使用‘：’】 d = %f\n", d)
	fmt.Println("声明全局变量：", a, b, c)

	//只能在函数体内声明
	g, h := 100, "你好"
	fmt.Println("':'声明变量只能在func函数体内实现", g, h)

	/*	//不能对g变量再次做初始化声明
		g := 400 // 语法就会异常 可以使用等号赋值，不能:=初始化
		fmt.Println("再次初始化g变量：", g, h)
	*/

	_, value := 7, 5 //实际上7的赋值被废弃，变量 _  不具备读特性
	//fmt.Println(_) //_变量的是读不出来的
	fmt.Println(value) //5

}
