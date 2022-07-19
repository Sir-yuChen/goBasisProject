package main

import (
	"fmt"
)

func main() {
	//第一种声明
	var test1 map[string]string
	//在使用map前，需要先make，make的作用就是给map分配数据空间
	test1 = make(map[string]string, 10)
	test1["one"] = "php"
	test1["two"] = "golang"
	test1["three"] = "java"
	fmt.Println("map声明方式一：", test1) //map[two:golang three:java one:php]

	//第二种声明
	test2 := make(map[string]string)
	test2["one"] = "php"
	test2["two"] = "golang"
	test2["three"] = "java"
	fmt.Println("map声明方式二：", test2) //map[one:php two:golang three:java]

	//第三种声明
	test3 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "java",
	}
	fmt.Println("map声明方式三：", test3) //map[one:php two:golang three:java]

	language := make(map[string]map[string]string)
	language["php"] = make(map[string]string, 2)
	language["php"]["id"] = "1"
	language["php"]["desc"] = "php是世界上最美的语言"
	language["golang"] = make(map[string]string, 2)
	language["golang"]["id"] = "2"
	language["golang"]["desc"] = "golang抗并发非常good"

	fmt.Println("map中的key依然是map结构：", language) //map[php:map[id:1 desc:php是世界上最美的语言] golang:map[id:2 desc:golang抗并发非常good]]

	//增删改查
	val, key := language["php"] //查找是否有php这个子元素
	if key {
		fmt.Println("map中key为‘php’的value值是：=====>", val)
	} else {
		fmt.Println("no")
	}

	language["php"]["id"] = "3"             //修改了php子元素的id值
	language["php"]["nickname"] = "这个一个key" //增加php元素里的nickname值
	delete(language, "php")                 //删除了php子元素
	fmt.Println("最终修改后的map集合：", language)
}
