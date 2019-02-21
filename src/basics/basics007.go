// basics007
/*
Go语言输出正弦函数（sin）图像
*/
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	/*图片大小*/
	const size = 300

	/*根据给定大小创建灰度图*/
	pic := image.NewGray(image.Rect(0, 0, size, size))

	/*遍历每个像素*/
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			/*填充为白色*/
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	/*从0到最大像素生成x坐标*/
	for x := 0; x < size; x++ {
		/*让sin的值的范围在0-2Pi之间*/
		s := float64(x) * 2 math.Pi / size
		
	}
}
