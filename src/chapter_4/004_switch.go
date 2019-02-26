// 004_switch
package main

import (
	"fmt"
)

func main() {
	var name = "Zhangsan"
	switch name {
	case "Lisi":
		fmt.Println("yes_1")
	case "Zhangsan", "Wangwu":
		fmt.Println("yes_2")
	case "Liuliu":
		fmt.Println("yes_3")
	default:
		fmt.Println("sorry")
	}
}
