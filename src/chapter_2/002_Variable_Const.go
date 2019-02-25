/*
第五节，Go语言常量和const关键字
相对于变量，常量是恒定不变的值，例如圆周率。
可以在编译时，对常量表达式进行计算求值，并在运行期使用该值，但计算结果不能修改。
*/
package main

import (
	"fmt"
)

const pi = 3.1415926
const a = 2.718281

const (
	pi_1 = 3.1415926
	a_2  = 2.718281
)

const size = 5

func main() {
	var arr [size]int
	fmt.Println(arr)
}
