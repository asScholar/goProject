// test003
/*
常量中的数据类型只可以是布尔型，数字型（整数型，浮点型和复数）和字符串型。
*/
package main

import (
	"fmt"
	"unsafe"
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a) /*常量可以用len(), cap(), unsafe.Sizeof()等函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过。*/
)

const (
	d = iota
	e = iota
	f = iota
	/*
		iota，特殊常量，可以认为是一个可以被编译器修改的常量。
		iota在const关键字出现时将被重置为0（const内部的第一行之前），const中每新增加一行常量声明将使iota计数一次。
	*/
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Println("the area is : %d", area)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
