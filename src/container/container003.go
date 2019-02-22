// container003
/*
Go语言append()为切片添加元素
*/
package main

import (
	"fmt"
)

func main() {
	var car []string

	/*添加一个元素*/
	car = append(car, "oldDriver")
	/*添加多个元素*/
	car = append(car, "Ice", "Sniper", "Monk")
	/*添加切片*/
	team := []string{"Pig", "Flyingcake", "Chicken"}
	car = append(car, team...)

	fmt.Println(car)
}
