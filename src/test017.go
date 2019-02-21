// test017
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
Go允许使用go语句开启一个新的运行期线程
*/
func main() {
	go say("world")
	say("hello")
}
