// process002
/*
Go语言for（循环结构）
*/
package main

import (
	"fmt"
)

func main() {
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}

	var i int
	for ; ; i++ {
		fmt.Println(i)
		if i > 10 {
			break
		}
	}
}
