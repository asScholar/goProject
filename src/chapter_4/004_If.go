// 004_If
package main

import (
	"fmt"
)

func getAge() int {
	return 11
}
func main() {
	var num int = 11
	if num > 10 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	if age := getAge(); age > 10 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
