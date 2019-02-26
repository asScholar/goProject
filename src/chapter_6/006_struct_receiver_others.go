/*
第四节，Go语言接收器（其它类型）
*/
package main

import (
	"fmt"
)

/*将int定义为MyInt类型*/
type MyInt int

/*为MyInt添加爱IsZero()方法*/
func (m MyInt) isZero() bool {
	return m == 0
}

func (m MyInt) Add(other int) int {
	return other + int(m)
}
func main() {
	var m MyInt
	fmt.Println(m.isZero())
	m = 1
	fmt.Println(m.Add(3))
}
