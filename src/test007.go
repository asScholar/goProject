// test007
package main

import (
	"fmt"
)

/*
全局变量
*/
var d int

/*
形式参数
*/
func sum(a, b int) int {
	return a + b
}

func main() {

	/*
		局部变量
	*/
	var a, b, c int
	a = 10
	b = 20
	c = a + b

	fmt.Println(a, b, c)
	fmt.Println(sum(a, b))
}
