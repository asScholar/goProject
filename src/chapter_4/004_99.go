// 004_99
package main

import (
	"fmt"
)

func main() {
	for x := 1; x <= 9; {
		for y := 1; y <= x; {
			fmt.Printf("%d x %d = %d\t", y, x, x*y)
			y++
		}
		fmt.Println()
		x++
	}
}
