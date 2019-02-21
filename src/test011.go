// test011
package main

import (
	"fmt"
)

func main() {
	/*创建切片*/
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(nums)

	/*打印原始切片*/
	fmt.Println(nums)

	/*打印子切片，从索引1（包含）到索引4（不包含）*/
	fmt.Println(nums[1:4])

	/*打印子切片，从索引3（包含）到索引4（不包含）*/
	fmt.Println(nums[:3])

	/*追加切片*/
	nums = append(nums, 10, 11)
	printSlice(nums)

	/*创建切片，nums1的容量是nums的两倍，接着拷贝切片*/
	nums1 := make([]int, len(nums), (cap(nums))*2)
	copy(nums1, nums)
	printSlice(nums1)
}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice=%v\n", len(x), cap(x), x)
}
