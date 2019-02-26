/*
Go语言的defer语句会将其后面跟随的语句进行延迟处理。
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("defer begin!")

	/*将defer放入延迟调用栈*/
	defer fmt.Println(1)
	defer fmt.Println(2)
	/*最后一个放入，位于栈顶，最先调用*/
	defer fmt.Println(3)

	fmt.Println("defer end")
}