package main

import "fmt"

type Rect struct {
	x      float64
	y      float64
	width  float64
	height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}
func main() {

	var rect1 = new(Rect)
	rect1.height = 3.0
	rect1.width = 7.21
	fmt.Println("the rect1: ", rect1.Area())

	//var a int = 12
	//b := "sunrisme"
	//var c int

	//var d *int

	//fmt.Println("a 变量地址：", &a)
	//fmt.Println("b 变量地址：", &b)
	//fmt.Println("c 的地址：", &c)

	//指针变量
	//var aPoint *int
	//aPoint = &a
	//fmt.Println("指针指向的内存地址中的值：", *aPoint)

	//fmt.Println("指针指向的内存地址中的值: ", d)
}
