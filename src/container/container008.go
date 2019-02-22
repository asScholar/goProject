// container008
/*
Go语言sync.Map（在并发环境中使用map）
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map

	/*将键值对保存到sync.Map*/
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egpyt", 200)

	fmt.Println(scene.Load("london"))
}
