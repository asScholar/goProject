/*
Go语言初始化内嵌结构体
*/
package main

import (
	"fmt"
)

/*汽车*/
type Car struct {
	Basics struct {
		Name  string
		Color string
	}
	Wheel
	Engine
}

/*车轮*/
type Wheel struct {
	Size int
}

/*引擎*/
type Engine struct {
	Power int
	Type  string
}

func main() {
	car := Car{
		Basics: struct {
			Name  string
			Color string
		}{
			Name:  "Audi",
			Color: "White",
		},
		/*初始化车轮*/
		Wheel: Wheel{
			Size: 18,
		},
		/**/
		Engine: Engine{
			Power: 143,
			Type:  "1.4T",
		},
	}
	fmt.Println(car)
}
