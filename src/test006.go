// test006
/*
func function_name([parameter list])[return_types]{
	function body
}
*/
package main

import (
	"fmt"
)

/*
多个输入参数，一个输出结果
*/
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/*
多个输入参数，多个输出结果
*/
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	var a int = 100
	var b int = 200
	var res int

	res = max(a, b)
	fmt.Println("the max is ", res)

	var str1 string = "Go!"
	var str2 string = "Hello "
	//var str11, str22 string

	str11, str22 := swap(str1, str2)

	fmt.Println(str11, str22)
}
