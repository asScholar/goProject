// container009
/*
Go语言list（列表）
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*初始化列表*/
	test1 := list.New() /*method 1*/
	//var test2 list.List /*method 2*/

	/*在列表中插入元素*/
	test1.PushBack("first")
	test1.PushBack(18)
	test1.PushFront("Hello")

	/*遍历列表 - 访问列表的每一个元素*/
	for i := test1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
