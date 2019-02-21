// test004
package main

import (
	"fmt"
)

func main() {
	var a int = 21
	var b int = 10
	var c int

	/*
		算数运算符
	*/
	c = a + b
	fmt.Println("第一行 – c 的值是 ", c)

	c = a - b
	fmt.Println("第二行 – c 的值是 ", c)

	c = a * b
	fmt.Println("第三行 – c 的值是 ", c)

	c = a / b
	fmt.Println("第四行 – c 的值是 ", c)

	c = a % b
	fmt.Println("第五行 – c 的值是 ", c)

	c++
	fmt.Println("第六行 – c 的值是 ", c)

	c--
	fmt.Println("第七行 – c 的值是 ", c)

	/*
		关系运算符
	*/
	if a == b {
		fmt.Println("第一行 – a is equal to b")
	} else {
		fmt.Println("第一行 – a is not equal to b")
	}

	if a > b {
		fmt.Println("第一行 – a is bigger than b")
	} else {
		fmt.Println("第一行 – a is less than b")
	}

	/*
		逻辑运算符
	*/
	var d bool = true
	var e bool = false

	if d && e {
		fmt.Println("both are true")
	}

	if d || e {
		fmt.Println("one is true")
	}

	/*
		位运算符
	*/

	var f uint = 60 /*60 = 0011 1100*/
	var g uint = 13 /*13 = 0000 1101*/
	// var h uint = 0

	fmt.Println("f & g = ", f&g)
	fmt.Println("f | g = ", f|g)
	fmt.Println("f ^ g = ", f^g)
	fmt.Println("f << 2 = ", f<<2)
	fmt.Println("f >> 2", f>>2)
}
