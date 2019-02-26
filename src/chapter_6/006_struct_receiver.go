/*
第四节，Go语言接收器(非指针型)
*/
package main

import (
	"fmt"
)

/*定义结构体*/
type Point struct {
	X int
	Y int
}

/*非指针接收器的加法*/
func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func main() {
	p1 := Point{1, 1}
	p2 := Point{2, 2}

	result := p1.Add(p2)

	fmt.Println(result)

}
