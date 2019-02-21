// test001

/*
定义了包名，必须在源码文件中非注释的第一行指明本文件属于哪个包。
Package main表示一个可独立执行的程序，每个Go应用程序都包含一个名为main的包。
*/
package main

/*
导入fmt包，fmt包实现了格式化I/O函数。
*/
import (
	"fmt"
)

/*
main函数，main函数是每一个可执行程序所必须包含的。
一般来说main函数是程序启动后第一个执行的函数（如果有init函数，则会首先执行init函数）。
*/
func main_() {
	/*
		标准输出函数。
	*/
	fmt.Println("Hello Go!")
}

/*
注意：
当标识符（包括常量，变量，类型，函数名，结构字段等等）以大写字母开头，那么该标识符所指对象可以被外部包的代码所使用，这被称为导出（就像面向对象语言中的public）；
如果标识符以小写字母开头，则表示对包外不可见，但是它们在整个包内部是可见且可用的（就像面向对象语言中的protected）。
*/