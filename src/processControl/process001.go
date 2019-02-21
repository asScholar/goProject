// process001
/*
Go语言if else（分支结构）
*/
package main

import (
	"fmt"
)

func sum(num1, num2 int) int {
	var result int
	result = num1 + num2
	return result
}

func main() {
	var num int = 11
	if num > 10 {
		fmt.Println(">10")
	} else {
		fmt.Println("<=10")
	}

	if res := sum(1, 2); res > 2 {
		fmt.Println(res)
	}
}
