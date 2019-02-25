// 002_Variable_enumeration
package main

import (
	"fmt"
)

/*声明芯片类型*/
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) getChip() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	fmt.Printf("%s %d\n", CPU, CPU)
}
