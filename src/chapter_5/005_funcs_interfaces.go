/*
函数和其他类型一样都属于“一等公民”，其他类型能够实现接口，函数也可以。
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

func main() {
	fmt.Println("Hello World!")
}
