// func001
/*
Go语言函数声明（函数定义）
*/
package main

import (
	"fmt"
)

/*普通函数的声明形式*/
func sum1(num1 int, num2 int) int {
	return num1 + num2
}

/*参数类型的简写*/
func sum2(num1, num2 int) int {
	return num1 + num2
}

/*函数的返回值*/
func sum_minutes(num1, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

/*带有变量名的返回值*/
func nameRetValues() (a, b int) {
	a = 1
	b = 2
	return
}

func main() {
	a, b := sum_minutes(1, 2)
	c, d := nameRetValues()
	fmt.Println(a, b, c, d)
}
