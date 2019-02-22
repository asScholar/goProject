// container007
/*
Go语言遍历map - （访问map中的每一个键值对）
*/
package main

import (
	"fmt"
)

func main() {
	scores := make(map[string]int)
	scores["Zhangsan"] = 80
	scores["Lisi"] = 79
	scores["Wangwu"] = 90
	scores["Liuliu"] = 12

	/*删除map指定元素*/
	delete(scores, "Liuliu")

	/*map遍历*/
	for k, v := range scores {
		fmt.Println(k, v)
	}

	/*复制map到切片*/
	var stus []string
	for k := range scores {
		stus = append(stus, k)
	}
	fmt.Println(stus)
}
