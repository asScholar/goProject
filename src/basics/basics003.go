// basics003
/*
Go语言多个变量同时赋值
*/
package main

import (
	"fmt"
)

func main() {

	/*数据交换，经典方式*/
	var a int = 100
	var b int = 200
	var t int

	t = a
	a = b
	b = t

	fmt.Println(a, b)

	/*数据交换，省内存方式*/
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(a, b)

	/*数据交换，Go语言方式*/
	a, b = b, a
	fmt.Println(a, b)
}
