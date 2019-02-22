// container004
/*
Go语言copy() - 切片复制（切片拷贝）
*/
package main

import (
	"fmt"
)

func main() {
	/*设置元素数量为1000*/
	const elementCount = 1000

	/*预分配足够多的元素切片*/
	srcData := make([]int, elementCount)

	/*为切片赋值*/
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	/*预分配足够多的元素切片*/
	copyData := make([]int, elementCount)
	fmt.Println(copyData[2:10])
	/*切片复制*/
	copy(copyData, srcData)

	fmt.Println(copyData[2:10])

	/*切片复制（指定区间）*/
	copy(copyData, srcData[3:8])
	fmt.Println(copyData[2:10])
}
