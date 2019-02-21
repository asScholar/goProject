// basics004
/*
Go语言匿名变量（没有名字的变量）
*/
package main

import (
	"fmt"
)

func getData() (int, int) {
	return 100, 200
}

func main() {
	var a, b int
	a, _ = getData()
	_, b = getData()
	fmt.Println(a, b)
}
