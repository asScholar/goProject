// process07
/*
Go语言break（跳出循环）
*/
package main

import (
	"fmt"
)

func main() {
OuterLoop:
	for x := 0; x < 2; x++ {
		for y := 0; y < 5; y++ {
			switch y {
			case 2:
				fmt.Println(x, y)
			case 3:
				fmt.Println(x, y)
				//break OuterLoop
				continue OuterLoop
			}
		}
	}
}
