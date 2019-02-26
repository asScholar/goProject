/*
使用SQL语言从数据库中获取数据时，可以对原始数据进行排序、分组和去重等。
对数据的操作进行多步骤的处理被称为链式处理。
*/
package main

import (
	"fmt"
	"strings"
)

/*字符串链式处理函数*/
func StringProcess(list []string, chain []func(string) string) {
	/*遍历每一个字符串*/
	for index, str := range list {
		/*第一个需要处理的字符串*/
		result := str

		/*遍历每一个处理链*/
		for _, proc := range chain {
			/*使用处理链对字符串进行处理，结果再存回*/
			result = proc(result)
		} //第一个字符串处理完毕，将其存回切片
		list[index] = result
	}
}

/*删除指定的前缀*/
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	/*准备待处理的字符串*/
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	/*准备处理链*/
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	/*使用处理链对字符串切片进行处理*/
	StringProcess(list, chain)
	for _, str := range list {
		fmt.Println(str)
	}
}
