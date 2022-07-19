package main

import "fmt"

/*
	slice: 切片  【对数组的抽象、动态数组】
		Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，
		功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大【扩大为原来切片容量的一倍】
	定义语法：
		var identifier []type
		无需长度
*/

func main() {
	/*
		使用make()函数来创建切片
			var slice1 []type = make([]type, len)
		简写为：
			slice1 := make([]type, len)
		也可以指定容量，其中capacity为可选参数
			make([]T, length, capacity)
		len 是数组的长度并且也是切片的初始长度
	*/
	/*
		s := []int{1, 2, 3}  直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
		s := arr[:]		初始化切片s,是数组arr的引用
		s := arr[startIndex:endIndex]		将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
		s := arr[startIndex:]		缺少endIndex时将表示一直到arr的最后一个元素
		s := arr[:endIndex]		缺少startIndex时将表示从arr的第一个元素开始
		s1 := s[startIndex:endIndex]		通过切片s初始化切片s1
		s :=make([]int,len,cap)		通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
	*/

	/*	len() 和 cap() 函数
		切片是可索引的，并且可以由 len() 方法获取长度
		切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少
	*/
	//示例：
	var num = make([]int, 3, 5) //创建一个分片 长度3 容量5 并且开辟出内存空间
	printSlice(num)             //执行结果：len=3 cap=5 slice=[0 0 0]

	//空切片
	var nullSlice []int
	printSlice(nullSlice)
	if nullSlice == nil {
		fmt.Println("nullSlice 切片为空")
		/*	执行结果：
			len=0 cap=0 slice=[]
			nullSlice 切片为空
		*/
	}
	fmt.Println("=====================>")

	//切片截取	【可以通过设置下限及上限来设置截取切片[lower-bound:upper-bound]】 包前不包后
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("原始切片	numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("截取切片	numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("截取默认下限 0	numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("截取默认上限为 len(s)	numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	fmt.Println("==========append() 和 copy() 函数===========")

	/*
		append() 和 copy() 函数
			如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来
	*/
	var small []int
	printSlice(small)

	/* 允许追加空切片 */
	small = append(small, 0)
	printSlice(small)

	/* 向切片添加一个元素 */
	small = append(small, 1)
	printSlice(small)

	/* 同时添加多个元素 */
	small = append(small, 2, 3, 4)
	printSlice(small)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	bigSlice := make([]int, len(small), (cap(small))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(bigSlice, small)
	printSlice(bigSlice)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
