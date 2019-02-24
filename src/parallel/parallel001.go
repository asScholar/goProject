// parallel001
/*
Go语言goroutine（轻量级线程）
*/
package main

import (
	"fmt"
	"time"
)

/*使用普通函数创建goroutine*/
func running() {
	var times int
	for {
		times++
		fmt.Println("tick out", times)

		/*间隔1秒*/
		time.Sleep(time.Second)
	}
}

func main() {

	/*并发执行程序*/
	go running()

	/*接收命令行输入，不做任何事，只用来证明可以并发执行*/
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

	/*使用匿名函数创建goroutine*/
	go func() {
		var times int
		for {
			times++
			fmt.Println("tick inner", times)
			time.Sleep(time.Second)
		}
	}()
}
