/*
第三节，Go语言指针
指针概念在Go语言中被拆分成两个概念
1. 类型指针，允许这个指针类型的数据进行修改。传递数据使用指针，而无须拷贝数据。类型指针不能进行偏移和运算。
2. 切片，由指向起始元素的原始指针、元素和容量组成。
*/
package main

import (
	"flag"
	"fmt"
)

func main() {
	/*认识指针地址和指针类型*/
	var cat int = 1
	var name string = "Zhangsan"

	fmt.Printf("%p %p\n", &cat, &name)

	/*从指针获取指针指向的值*/
	var house = "Malibu Point 10880, 90265"
	ptr := &house
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("address: %p\n", ptr)
	value := *ptr
	fmt.Printf("%s, %s\n", house, value)

	/*使用指针修改值*/
	*ptr = "Xueyuan Nan 37, Haidian Beijing"
	fmt.Printf("%s, %s\n", house, value)

	/*使用指针变量获取命令行的输入信息.执行命令：go run 002_Variable_Pointer.go --mode=fast*/
	var mode = flag.String("mode", "", "process mode") //定义命令行参数
	flag.Parse()                                       //解析命令行参数
	fmt.Println(*mode)                                 //输出命令行参数

	/*创建指针的另一种方法 - new()函数*/
	dog := new(string)
	*dog = "Damao"
	fmt.Println(*dog)
}
