/*第二节，Go语言变量类型
Go语言中有丰富的数据类型，除了基本的类型的整型、浮点型、布尔型、字符串之外，还有切片、结构体、函数、map、通道等。
*/
/*
整型分为以下两大类
1. 按长度分为：int8,int16,int32,int64
2. 还有对应的无符号整型：uint8,uint16,uint32,uint64
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Go语言支持两种浮点型数：float32,float64
	*/
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	/*布尔型数据在Go语言中以bool类型进行声明，布尔型数据只有true和false两个值,默认值为false。*/
	var n bool
	fmt.Println(n)

	/*字符串在Go语言中以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32等）一样。*/
	str := "Hello Go!"
	fmt.Println(str)

	/*Go语言的字符串常见转义符*/
	fmt.Println("\"Hello everyone!\"")

	/*多行字符串的使用,注意这里使用的不是单引号，而是键盘1左边的键.多行字符串一般用于内嵌源码和内嵌数据等.*/
	const mulStr = `
	frist
	second
	third
	`
	fmt.Println(mulStr)

	/*
		Go语言的字符有以下两种：
		1. unit8类型，或者叫byte型，代表ASCII码的一个字符。
		2. rune类型，代表UTF-8字符。当需要处理中文、日文或者其他复合字符时，需要使用rune类型。rune实际上是一个int32.
	*/
	var test_1 byte = 'a'
	var test_2 rune = '你'
	fmt.Printf("%d %T\n", test_1, test_1)
	fmt.Printf("%d %T\n", test_2, test_2)

	/*类型转换*/
	var pi = math.Pi
	fmt.Println(pi)
	fmt.Println(int(pi))

}
