// test008
package main

import (
	"fmt"
)

func main() {
	var nums5 [5]int
	var i, j, k int

	for i = 0; i < 5; i++ {
		nums5[i] = i
	}
	for j = 0; j < 5; j++ {
		fmt.Println(nums5[j])
	}

	var nums = [...]int{1, 2, 3, 4, 5}
	for k = 0; k < 5; k++ {
		fmt.Println(nums[k])
	}
}
