// func003
/*
了解Go 语言的参数值传递
*/
package main

import (
	"fmt"
)

type Data struct {
	complax  []int
	instance InterData
	ptr      *InterData
}
type InterData struct {
	a int
}

func passByValue(inFunc Data) Data {
	fmt.Printf("inFunc value: %+v\n", inFunc)
	fmt.Printf("inFunc ptr:%p\n", &inFunc)
	return inFunc
}

func main() {
	in := Data{
		complax:  []int{1, 2, 3},
		instance: InterData{5},
		ptr:      &InterData{1},
	}

	fmt.Printf("in value: %+v\n", in)
	fmt.Printf("in ptr: %p\n", &in)
	out := passByValue(in)
	fmt.Printf("out value: %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)
}
