// container002
/*
Go语言切片详解
*/
package main

import (
	"fmt"
)

func main() {
	/*从数组生成新的切片*/
	var nums = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(nums, nums[1:5])

	/*从指定范围中生成切片*/
	/*全部元素*/
	fmt.Println(nums[:])
	/*区间*/
	fmt.Println(nums[3:8])
	/*中间到尾部的所有元素*/
	fmt.Println(nums[4:])
	/*开头到中间的所有元素*/
	fmt.Println(nums[:8])

	/*重置切片，清空拥有的元素*/
	fmt.Println(nums[0:0])

	/*直接声明新的切片*/
	var strList []string
	var numList []int

	/*声明一个空切片*/
	var numListEmpty = []int{}
	fmt.Println(strList, numList, numListEmpty)
	fmt.Println(len(strList), len(numList), len(numListEmpty))

	/*切片判空*/
	fmt.Println(strList == nil)

	/*使用make()函数构造切片*/
	/*
		make([]T, size, cap)
		T:切片的元素类型
		size:为这个类型分配多少个元素
		cap:预分配的元素数量，这个值设定后不影响size，只是能提前分配空间，降低多次分配空间的性能问题
	*/
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(len(a), len(b)) /*a，b均是预分配2个元素的切片，只是b的内部存储空间已经分配链10个，但实际使用链2个元素*/
}
