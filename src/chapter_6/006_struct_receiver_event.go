/*
第四节，Go语言接收器（事件系统）
事件系统可以将事件派发者与事件处理者解耦。
*/
package main

import (
	"fmt"
)

var eventByName = make(map[string][]func(interface{}))

func RegisterEvent(name string, callback func(interface{})) {
	list := eventByName[name]
	list = append(list, callback)
	eventByName[name] = list
}

func CallEvent(name string,param interface{}){
	list:=eventByName[name]
	for_,callback:=range list{
		callback(param)
	}
}

func main() {
	fmt.Println("Hello World!")
}
