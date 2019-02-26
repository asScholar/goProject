/*
第四节，Go语言接收器(指针型)
*/
package main

import (
	"fmt"
)

/*定义属性结构体*/
type Property struct {
	value int //属性值
}

/*设置属性值*/
func (p *Property) SetValue(v int) {
	p.value = v //修改p的成员变量
}

/*取属性值*/
func (p *Property) GetValue() int {
	return p.value
}

func main() {
	p := new(Property)
	p.SetValue(1001)
	fmt.Println(p.GetValue())
}
