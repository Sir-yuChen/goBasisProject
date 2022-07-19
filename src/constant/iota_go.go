package main

import "fmt"

func main() {
	//在 golang 中，一个方便的习惯就是使用iota标示符，它简化了常量用于增长数字的定义，给以上相同的值以准确的分类

	//自增1
	const (
		CategoryBooks    = iota // 0
		CategoryHealth          // 1
		CategoryClothing        // 2
	)
	fmt.Println("iota 自增：", CategoryBooks, CategoryHealth, CategoryClothing)
	/*
		iota和表达式
			iota可以做更多事情，而不仅仅是 increment。更精确地说，iota总是用于 increment，但是它可以用于表达式，在常量中的存储结果值
	*/
	type Allergen int
	const (
		IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
		IgChocolate                         // 1 << 1 which is 00000010
		IgNuts                              // 1 << 2 which is 00000100
		IgStrawberries                      // 1 << 3 which is 00001000
		IgShellfish                         // 1 << 4 which is 00010000
	)
	//这个工作是因为当你在一个const组中仅仅有一个标示符在一行的时候，它将使用增长的iota取得前面的表达式并且再运用它，。在 Go 语言的spec中， 这就是所谓的隐性重复最后一个非空的表达式列表.
	fmt.Println("iota和表达式：", IgEggs, IgChocolate, IgNuts, IgStrawberries, IgShellfish)

	const (
		Apple, Banana     = iota + 1, iota + 2
		Cherimoya, Durian //下一行增长，而不是立即取得它的引用。
		Elderberry, Fig
	)
	fmt.Println("iota和表达式 多变量相同行：", Apple, Banana)
	fmt.Println("iota和表达式 多变量相同行：", Cherimoya, Durian)
	fmt.Println("iota和表达式 多变量相同行：", Elderberry, Fig)

}
