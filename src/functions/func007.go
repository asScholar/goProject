// func007
/*
Go语言闭包（Closure） - 引用链外部变量的匿名函数
*/
package main

import (
	"fmt"
)

/*提供一个值，每次调用函数指定对值进行累加*/
func Accumulate(value int) func() int {
	/*返回一个闭包*/
	return func() int {
		value++
		return value
	}
}
func main() {
	/*创建一个累加器，初始值为1*/
	acculator := Accumulate(1)

	fmt.Println(acculator())
	fmt.Println(acculator())
	fmt.Printf("%p\n", acculator)

	acculator2 := Accumulate(10)

	fmt.Println(acculator2())
	fmt.Println(acculator2())
	fmt.Printf("%p\n", acculator2)

	/*在闭包内部修改引用的变量*/
	str := "hello go"
	foo := func() {
		str = "hello world"
		fmt.Println(str)
	}
	foo()
	fmt.Println(str)
}
