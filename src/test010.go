// test10
/*
type struct_variable_type struct {
	member definition;
	member definition;
	......
	member definition;
}
*/
package main

import (
	"fmt"
)

/*
定义结构体
*/
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

/*
结构体作函数参数
*/
func printBook(book Books) {
	fmt.Println(book)
}

func main() {
	/*
		结构体赋值
	*/
	var book = Books{
		title:   "Go learning",
		author:  "Zhu Jinbin",
		subject: "Software",
		book_id: 1001,
	}
	fmt.Println(book)
	fmt.Println(book.author)

	printBook(book)
}
