/*
第三节，Go语言方法和接收器
Go语言中的方法是一种作用于特定类型变量的函数，这种特定类型变量叫做接收器。
如果将特定类型理解为结构体或“类”时，接收器的概念就类似于其他语言中的this或self。
在Go语言中接收器可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
以下代码使用背包作为“对象”，将物品放入背包的过程称为“方法”，通过面向过程的方式和Go语言中结构体的方式来理解“方法”的概念。
*/
package main

import (
	"fmt"
)

/*背包*/
type Bag struct {
	items []int
}

/*将物品放入背包*/
func Insert(bag *Bag, itemid int) {
	bag.items = append(bag.items, itemid)
}

/*接收器的方式实现结构体操作*/
func (bag *Bag) Insert_Receiver(itemid int) {
	bag.items = append(bag.items, itemid)
}
func main() {
	bag := new(Bag)
	// Insert(bag, 1001)
	// Insert(bag, 1002)
	// Insert(bag, 1003)
	bag.Insert_Receiver(1001)
	bag.Insert_Receiver(1002)
	bag.Insert_Receiver(1005)
	fmt.Println(bag)
}
