// basics002
/*
Go语言变量的初始化
*/
package main

import (
	"fmt"
)

/*变量初始化的标准格式*/
var hp_1 int = 100

/*编译器推导类型的格式*/
var hp_2 = 100

/*短变量声明并初始化*/
hp_3 := 100

func main() {
	fmt.Println("Hello World!")
}
