// process005
/*
Go语言匿名函数 - 没有函数名字的函数
*/
package main

import (
	"flag"
	"fmt"
)

/*匿名函数作回调函数*/
func visist(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {
	/*在定义时调用匿名函数*/
	func(data int) {
		fmt.Println("hello", data)
	}(100)

	/*将匿名函数赋值给变量*/
	f := func(data int) {
		fmt.Println("hello", data)
	}
	f(100)

	/*匿名函数作回调函数*/
	visist([]int{1, 2, 3, 4, 5}, func(v int) {
		fmt.Println(v)
	})

	/*使用匿名函数实现操作的封装*/ //不懂
	var skillParam = flag.String("skill", "", "skill to perform")
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}

}
