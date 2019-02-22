// func006
/*
Go语言函数类型实现接口 - 把函数作为接口来调用
*/
package main

import (
	"fmt"
)

/*调用器接口*/
type Invoker interface {
	/*需要实现一个Call方法*/
	Call(interface{})
}

/*结构体类型*/
type Struct struct{}

/*实现Invoker的Call*/
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

/*把函数定义为类型*/
type FuncCaller func(interface{})

func main() {
	fmt.Println("Hello World!")
}
