// test015
package main

import (
	"fmt"
)

func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	/*强制类型转换*/
	mean = float32(sum) / float32(count)
	fmt.Printf("%f\n", mean)
}
