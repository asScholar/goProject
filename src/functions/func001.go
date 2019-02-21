// func001
/*
Go语言函数声明（函数定义）
*/
package main

import (
	"fmt"
)

/*普通函数的声明形式*/
func foo(a int, b string) {}

/*参数类型的简写*/
func sum(num1, num2 int) int {
	return num1 + num2
}

/*函数的返回值*/
func connectToNetwork() (int, int) {
	return 1, 2
}

/*带有变量名的返回值*/
func nameRetValues() (a, b int) {
	a = 1
	b = 2
	return
}

func main() {
	a, b := connectToNetwork()
	c, d := nameRetValues()
	fmt.Println(a, b, c, d)
}
