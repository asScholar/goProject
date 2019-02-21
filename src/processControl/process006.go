// process006
/*

 */
package main

import (
	"fmt"
)

func main() {
	/*goto跳出循环*/
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				goto breakHere
			}
		}
	}
	return
breakHere:
	fmt.Println("done")
}
