/*
第一节，Go语言数组和切片详解
数组（Array）是一段固定长度的连续内存区域。
切片是一个拥有相同类型元素的可变长度的序列。Go语言切片的内部结构包含地址、大小和容量。切片一般用于快速地操作一块数据集合。
*/
package main

import (
	"fmt"
)

func main() {
	/*数组声明*/
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "num"
	fmt.Println(team)

	/*数组初始化*/
	var cities = [...]string{"Beijing", "Shanghai", "Guangzhou", "Shenzhen", "Linyi", "Qingdao", "Jinan"}

	/*遍历数组元素*/
	for k, v := range cities {
		fmt.Println(k, v)
	}

	/*从数组或切片生成新的切片*/
	//从指定区间生成切片
	fmt.Println(cities[2:5])
	fmt.Println(cities[3:])
	fmt.Println(cities[:5])

	//表示原有的切片
	fmt.Println(cities[:])

	//重置切片，清空所有元素
	fmt.Println(cities[0:0])

	/*直接声明新的切片*/
	var strList []string
	var numList []int
	var numListEmpty = []int{}
	fmt.Println(strList, numList, numListEmpty)
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	fmt.Println(strList == nil) //切片判空

	/*使用make()函数构造切片*/
	students := make([]string, 2, 10) //2表示预分配2个元素，10表示其内部存储空间已经分配了10个
	fmt.Println(len(students))

	/*Go语言append()为切片添加元素*/
	var locations = cities[0:0]
	fmt.Println(locations)
	locations = append(locations, "Tianjin")
	locations = append(locations, "Xuzhou", "Nanjing")
	locations = append(locations, cities[:]...)
	fmt.Println(locations)

	/*切片复制*/
	fmt.Println("locations: ", locations)
	fmt.Println("cities: ", cities[:])
	copy(locations, cities[:])
	fmt.Println(locations)

	/*切片删除*/
	index := 2
	locations = append(locations[:index], locations[index+1:]...)
	fmt.Println(locations)
}
