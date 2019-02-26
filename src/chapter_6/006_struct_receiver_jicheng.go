/*
第四节，Go语言接收器（继承）
面向对象思想中，实现对象关系需要使用“继承”特性。例如，人类不能飞行，鸟类可以飞行。人类和鸟类都可以继承自可行走类，但只有鸟类继承飞行类。
Go语言的结构体内嵌特性就是一种组合特性，使用组合特性可以快速构建对象的不同特性。
*/
package main

import (
	"fmt"
)

/*可飞行的*/
type Flying struct{}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

/*可行走的*/
type Walkable struct{}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

/*人类，可行走*/
type Human struct {
	Walkable
}

/*鸟类，可行走、可飞行*/
type Bird struct {
	Walkable
	Flying
}

func main() {
	bird := new(Bird)
	fmt.Println("Bird: ")
	bird.Fly()
	bird.Walk()

	human := new(Human)
	fmt.Println("Human: ")
	human.Walk()
}
