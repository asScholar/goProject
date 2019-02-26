/*
Go语言支持匿名函数，即在需要使用函数时，再定义函数，匿名函数往往以变量方式被传递。
*/
package main

import (
	"fmt"
)

func main() {
	/*1. 在定义时调用匿名函数*/
	func(name string) {
		fmt.Println("Hello ", name)
	}("Zhangsan")

	/*2. 将匿名函数赋值给变量*/
	f := func(num1, num2 int) int {
		return num1 + num2
	}
	fmt.Println(f(1, 2))

	/*3. 匿名函数作回调函数*/
	visit([]int{1, 2, 3, 4, 5}, func(v int) {
		fmt.Println(v)
	})

	/*4. 使用匿名函数实现操作封装*/
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	var figure = "fly"
	switch figure {
	case "fire":
		skill["fire"]()
	case "fly":
		skill["fly"]()
	case "run":
		skill["run"]()
	default:
		fmt.Println("others")
	}
}

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
