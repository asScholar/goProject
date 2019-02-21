// test014
package main

import (
	"fmt"
)

/*
递归函数
*/
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	var num int = 15
	fmt.Println(Factorial(uint64(num))) /*强制类型转换*/
}
