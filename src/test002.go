// test002
package main

import (
	"fmt"
)

/*
类型相同的多个变量的声明。
局部变量声明之后必须调用，全局变量没有这个要求。
*/
var x, y int

/*
因式分解关键字的写法一般用于声明全局变量。
*/
var (
	a int
	b bool
)

/*
声明变量的同时进行初始化。
*/
var c, d int = 1, 2
var e, f = 123, "hello"

func main_() {
	g, h := 123, "Go!"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
