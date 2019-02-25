/*
第三节，Go语言list（列表）
列表是一种非连续存储的容器，由多个节点组成，节点通过一些变量记录彼此之间的关系。
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*初始化列表*/
	var names list.List
	//str := list.New()

	/*插入元素*/
	names.PushBack("Zhangsan")
	names.PushBack("Lisi")
	names.PushFront("Liuliu")
	names.PushBack("Wangsu")
	for name := names.Front(); name != nil; name = name.Next() {
		fmt.Println(name.Value)
	}

	names.Remove(names.Front())
	for name := names.Front(); name != nil; name = name.Next() {
		fmt.Println(name.Value)
	}

}
