/*
第四节，Go语言接收器
一个结构体的方法（class.Do)的参数和一个普通函数(funcDo)的参数完全一致。
*/
package main

import (
	"fmt"
)

type class struct{}

func (c *class) Do(v int) {
	fmt.Println("call method do: ", v)
}

func funcDo(v int) {
	fmt.Println("call function do: ", v)
}

func main() {
	/*声明一个回调函数*/
	var delegate func(int)

	c := new(class)
	delegate = c.Do
	delegate(100)

	delegate = funcDo
	delegate(100)
}
