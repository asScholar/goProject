/*
闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，
在闭包中可以继续使用这个自由变量。
*/
package main

import (
	"fmt"
)

func main() {
	/*1. 什么是闭包*/
	str := "Hello World!"
	foo := func() { //变量str定义在匿名函数之外，但str在匿名函数中被引用并被修改，此时便形成链闭包。
		str = "Hello Go!"
	}
	foo()

	fmt.Println("Hello World!")

	/*2. 闭包的记忆效应*/
	acculator := Accumulate(1)
	fmt.Println(acculator())
	fmt.Println(acculator())
	fmt.Println(acculator())
	fmt.Println(&acculator)
	acculator_2 := Accumulate(10)
	fmt.Println(acculator_2())
	fmt.Println(acculator_2())
	fmt.Println(acculator_2())
	fmt.Println(&acculator_2)

	/*3. 闭包实现生成器*/
	generator := playerGen("high noon")
	name, hp := generator()
	fmt.Println(name, hp)
}

/*提供一个值，每次调用函数就会对指定的值累加1*/
func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

/*创建一个玩家生成器，输入名称，输出生成器*/
func playerGen(name string) func() (string, int) {
	hp := 100
	return func() (string, int) {
		return name, hp
	}
}
