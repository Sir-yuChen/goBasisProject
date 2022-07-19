package main

import "fmt"

/*
	interface{}类型
		Golang的语言中提供了断言的功能。
		golang中的所有程序都实现了interface{}的接口，这意味着，所有的类型如string,int,int64甚至是自定义的struct类型都就此拥有了interface{}的接口，
		这种做法和java中的Object类型比较类似。那么在一个数据通过func funcName(interface{})的方式传进来的时候，也就意味着这个参数被自动的转为interface{}的类型。
*/

/*
//断言失败 cannot convert a (variable of type interface{}) to type string: need type assertion
func interFunc(a interface{}) string {
	return string(a)
}*/
func funcName(a interface{}) string {
	value, ok := a.(string)
	if !ok {
		fmt.Println("It is not ok for type string")
		return ""
	}
	fmt.Println("The value is ", value)

	return value
}

func main() {
	//s := interFunc("lll")
	//fmt.Println(s)
	/*
		执行结果：
			cannot convert a (variable of type interface{}) to type string: need type assertion
		此时，意味着整个转化的过程需要类型断言
	*/
	var a interface{}
	//fmt.Println("直接使用断言：", a.(string))	//直接使用断言 a.(string)	如果断言失败一般会导致panic的发生。所以为了防止panic的发生，我们需要在断言前进行一定的判断
	funcName(a)

}
