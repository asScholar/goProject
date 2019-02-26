// 005_funcs
package main

import (
	"fmt"
)

/*1. 函数声明*/
func foo(a int, b string) {}

func sum(num1, num2 int) int {
	return num1 + num2
}

/*2. 函数返回值*/
func typedTwoValues() (int, int) {
	return 1, 2
}
func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return
}

func main() {
	/*3. 函数调用*/
	a, b := namedRetValues()
	fmt.Println(a, b)

	/*4. 函数变量*/
	var f func(int, int) int
	f = sum
	fmt.Println(f(a, b))
}
