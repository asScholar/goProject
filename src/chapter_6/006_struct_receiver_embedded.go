// 006_struct_receiver_embedded
package main

import (
	"fmt"
)

/*基础颜色*/
type BasicColor struct {
	R, G, B float32
}

/*完整颜色定义*/
type Color struct {
	//Basic BasicColor //将基本颜色作为成员
	BasicColor
	Alpha float32 //透明度
}

func main() {
	var color Color

	color.R = 1
	color.G = 1
	color.B = 0

	color.Alpha = 1

	fmt.Println(color)
}
