package main

import (
	_ "../InitLib1" //注意导包，不要以 ‘/’结尾  ‘_’表示，加载该包，但是未使用包中的自定义方法，可能只需要包中init方法
	//_ "../InitLib2"
	"fmt"
)

//Go 函数可以返回多个值
func swap(x, y string) (string, string) {
	return x, y
}
func init() {
	fmt.Println("main=====>init")
}

func main() {
	fmt.Println("main 入口===》执行")

	a, b := swap("Holle ", "Go")
	fmt.Print("GO 函数可以返回多个值：", a, b)

	//init函数   init 函数可在package main中，可在其他package中，可在同一个package中出现多次。
	/*
			golang里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）。这两个函数在定义时不能有任何的参数和返回值。
			虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个package中每个文件只写一个init函数。
			go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。
			程序的初始化和执行都起始于main包。

		如果main包还导入了其它的包，那么就会在编译时将它们依次导入。有时一个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到fmt包，但它只会被导入一次，因为没有必要导入多次）。

		当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依次类推。

		等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在的话），最后执行main函数。
	*/

	/*
		GO导包和init()/main()函数执行顺序
			1. 在main中导入  InitLib1 InitLib2 两个包，执行结果：
				D:\Go\GoProject\basisProject\bin\go_build_func_go_go.exe
				IntLib1=====>init
				IntLib2=====>init
				main=====>init
				main 入口===》执行
				GO 函数可以返回多个值：Holle Go
				Process finished with the exit code 0
			从执行结果来看，go将所有的包一次性全部导入，导包的同时执行init保留函数 最后执行main入口init和main函数

			2. main中导入InitLib1 在InitLib1中导入InitLib2执行结果:
				IntLib2=====>init
				IntLib1=====>init
				main=====>init
				main 入口===》执行
				GO 函数可以返回多个值：Holle Go
				Process finished with the exit code 0
			从执行结果来看，go加载导包，时发现有多层包结构，会先找到最后一级并执行，并执行最后一级包init保留函数，依次向前推导包同时执行init函数，最后执行main入口init和main函数

		注意：相同的包结构，只会导入一次
	*/

}
