// test009
package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var ip *int

	ip = &a

	fmt.Println("变量地址：", &a, ip)
	fmt.Println(a, *ip)
}
