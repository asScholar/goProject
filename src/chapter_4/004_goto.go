// 004_goto
package main

import (
	"fmt"
)

func main() {
	var num = 9
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				num -= y
				goto breakHere
			}
		}
	}

breakHere:
	fmt.Println(num)
}
