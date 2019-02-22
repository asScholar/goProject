// container006
/*
Go语言map（Go语言映射）
*/
package main

import (
	"fmt"
)

func main() {
	scores := make(map[string]int)
	scores["Zhangsan"] = 89
	scores["Lisi"] = 90
	fmt.Println(scores["Lisi"])
}
