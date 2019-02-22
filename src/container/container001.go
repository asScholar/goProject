// container001
/*
Go语言数组详解
*/
package main

import (
	"fmt"
)

func main() {
	/*Go语言数组的声明*/
	var team [3]string
	team[0] = "China"

	/*Go语言数组的初始化*/
	var students = [...]string{"Zhangsan", "Lisi", "Wangwu"}

	/*遍历数组 - 访问每一个数组元素*/
	for k, v := range students {
		fmt.Println(k, v)
	}
}
