// 004_For
package main

import (
	"fmt"
)

func main() {
	var index int = 5
	for {
		if index < 0 {
			return
		} else {
			fmt.Println(index)
			index--
		}
	}
}
