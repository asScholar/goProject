// process004
/*
Go语言for range（键值循环）
*/
package main

import (
	"fmt"
)

func main() {
	/*遍历数组，切片 - 获得索引和元素*/
	for key, value := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("key: %d value: %d\n", key, value)
	}

	/*遍历字符串 - 获得字符*/
	var str = "Hello Go!"
	for key, value := range str {
		fmt.Printf("key: %d value: %c\n", key, value)
	}

	/*遍历map - 获得map的键和值*/
	m := map[string]int{
		"hello": 100,
		"go":    200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	/*遍历通道（channel）- 接收通道数据*/
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for value := range c {
		fmt.Println(value)
	}
}
