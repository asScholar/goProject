/*
第二节，Go语言map（映射）
Go语言提供的映射关系容器为map，map使用三列表（hash）实现。
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	/*声明map*/
	scene := make(map[string]int)

	/*添加元素*/
	scene["route"] = 77
	scene["brazil"] = 90
	scene["china"] = 78

	/*遍历、访问map*/
	for k, v := range scene {
		fmt.Println(k, v)
	}

	/*删除元素*/
	delete(scene, "brazil")
	for k, v := range scene {
		fmt.Println(k, v)
	}
}
