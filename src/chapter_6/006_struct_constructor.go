/*
第二节，Go语言类型或结构体没有构造函数，结构体的初始化过程可以使用函数封装实现。
如果使用结构体描述实例的特性，那么根据实例（以猫为例）的颜色和名字可以有不同种类的猫。
那么不同颜色和名字就是结构体的字段，同时可以使用颜色和名字构造不同种类猫的实例，这个过程可以参考如下代码。
*/
package main

import (
	"fmt"
)

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

func NewCat(name, color string) *Cat {
	return &Cat{
		Name:  name,
		Color: color,
	}
}

func main() {

	cat_1 := NewCatByName("Tom")
	cat_1.Color = "black"

	cat_2 := NewCat("Jerry", "white")
	fmt.Println(cat_1.Name, cat_1.Color)
	fmt.Println(cat_2.Name, cat_2.Color)
}
