/*
第一节，Go语言结构体(声明+实例化)
Go语言的关键字type可以将各种基本类型定义为自定义类型，基本类型包括整型、字符串、布尔等。
结构体是一种复合的基本类型，通过type定义为自定义类型后，使结构体便于使用。
*/
package main

import (
	"fmt"
)

/*1. 声明结构体*/
type Point struct {
	X int
	Y int
}

func main() {
	/*2. 实例化结构体(5种方法)*/
	var p_1 Point
	p_1.X = 10
	p_1.Y = 20

	p_2 := new(Point)
	p_2.X = 11
	p_2.Y = 21

	p_3 := &Point{}
	p_3.X = 12
	p_3.Y = 22

	p_4 := Point{
		X: 13,
		Y: 23,
	}

	p_5 := Point{
		14,
		24,
	}
	fmt.Println(p_1.X, p_2.X, p_3.X, p_4.X, p_5.X)

	/*3. 匿名结构体*/
	msg := &struct { //声明部分
		id   int
		data string
	}{ //值初始化部分
		1024,
		"Hello",
	}
	printMsgType(msg)

}

/*打印消息类型，传入匿名结构体*/
func printMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n", msg)
	fmt.Println(msg.data)
}
