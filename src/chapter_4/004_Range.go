// 004_Range
package main

import (
	"fmt"
)

func main() {
	/*遍历数组、切片*/
	for k, v := range []string{"Zhangsan", "Lisi", "Wangwu", "Liuliu"} {
		fmt.Println(k, v)
	}

	/*遍历字符串*/
	str := "Hello Go!"
	for k, v := range str {
		fmt.Println(k, v)
	}

	/*遍历map*/
	scores := map[string]int{
		"Zhangsan": 89,
		"Lisi":     78,
	}
	for k, v := range scores {
		fmt.Println(k, v)
	}

	/*遍历channel - 接收数据*/
	infos := make(chan int)
	go func() {
		infos <- 1
		infos <- 2
		infos <- 3
		close(infos)
	}()

	for v := range infos {
		fmt.Println(v)
	}
}
