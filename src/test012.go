// test012
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	/*range用于切片*/
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}

	/*range用于map*/
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s --> %s\n", k, v)
	}
}
