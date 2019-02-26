/*
所谓可变参数，是指指数数量不固定的函数形式。
*/
package main

import (
	"bytes"
	"fmt"
)

/*定义一个函数，参数数量为0-n，类型约束为字符串*/
func joinStrings(slist ...string) string {
	/*定义一个字节缓冲，快速地连接字符串*/
	var b bytes.Buffer
	/*遍历可变参数列表slist，类型为[]string*/
	for_,s:=range slist{
		/*将遍历出的字符串连续写入字节数组*/
		b.WriteString(s)
	}
	/*将连接好的字节数组转换为字符串数组并输出*/
	return b.String()
}

func main() {
	fmt.Println(joinStrings("pig","and","cat"))
}
