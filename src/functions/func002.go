// func002
/*
Go语言将秒转换为具体的时间
*/
package main

import (
	"fmt"
)

const (
	sedondPerMinute = 60
	secondPerHour   = sedondPerMinute * 60
	secondPerDay    = secondPerHour * 24
)

func resolveTime(seconds int) (day, hour, minute int) {
	day = seconds / secondPerDay
	hour = seconds / secondPerHour
	minute = seconds / sedondPerMinute
	return
}
func main() {
	fmt.Println(resolveTime(1000))
}
