/*
Go语言中传入和返回参数在调用和返回时都使用值传递，
这里需要注意的是指针、切片和map等引用型对象指向的内容在参数传递过程中是不会进行复制的，
而是将指针进行复制，类似于创建一次引用。
*/
package main

import (
	"fmt"
)

type Data struct {
	complax  []int
	instance InnerData
	ptr      *InnerData
}
type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {
	fmt.Printf("形参 value: %+v\n", inFunc)
	fmt.Printf("形参 ptr: %p\n", &inFunc)
	return inFunc
}

func main() {
	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}
	fmt.Printf("实参 value: %+v\n", in)
	fmt.Printf("实参 ptr: %p\n", &in)
	out := passByValue(in)
	fmt.Printf("结果 value: %+v\n", out)
	fmt.Printf("结果 ptr: %p\n", &out)
}
