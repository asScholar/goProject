// process005
/*
Go语言switch case语句
*/
package main

import (
	"fmt"
)

func main() {
	strs := "Hello"
	//name := "Zhangsan"
	switch strs {
	case "yes":
		fmt.Println(1)
	case "Hello Zhangsan":
		fmt.Println(2)
	case "Hello", "Zhangsan": /*匹配其中一个就是成功*/
		fmt.Println(3)
	default:
		fmt.Println(0)
	}

	num := 10
	switch {
	case num == 4:
		fmt.Println("yes")
	case num > 11:
		fmt.Println("no")
	case num > 5 && num < 12:
		fmt.Println("Good!")
	default:
		fmt.Println(0)
	}
}
