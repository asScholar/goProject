/*第一节，Go语言变量声明、初始化
Go语言的每一个变量都拥有自己的类型，必须经过声明才能开始用。
*/
package main

import (
	"fmt"
)

/*
变量声明，使用关键字var(标准格式)
*/
var a int
var b string
var c []float32
var d func() bool //声明一个返回值为布尔类型的函数变量，这种形式一般用于回调函数，即将函数以变量的形式保存下来，在需要的时候重新调用这个函数
var e struct {
	x int
}

/*
变量声明，使用关键字var(批量格式)
*/
var (
	f int
	g string
	h []float32
	i func() bool
	j struct {
		x int
	}
)

/*变量初始化，标准格式*/
var hp_1 int = 100

/*变量初始化，编译器推导类型的格式*/
var hp_2 = 101

func GetData() (int, int) {
	return 100, 200
}

func main() {
	/*短变量声明并初始化,这种方式声明的变量，一旦申明必须调用*/
	hp_3 := 102

	/*使用Go语言进行进行数据交换*/
	hp_1, hp_2 = hp_2, hp_1

	/*匿名变量*/
	hp_3, _ = GetData()

	fmt.Println(hp_1, hp_2, hp_3)
}
