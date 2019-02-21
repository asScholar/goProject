// test016
package main

import (
	"fmt"
)

/*
定义接口
type interface_name interface{
	method_name1 [return_type]
	method_name2 [return_type]
	method_name3 [return_type]
	method_name4 [return_type]
	......
}

定义结构体
type struct_name struct {

}
实现接口方法
func (struct_name_variable struct_name) method_name1()[return_type]{

}
*/

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("Nokia Phone")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("IPhone")
}
func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
