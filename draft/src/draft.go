package main

import "fmt"

func main() {
	var a int = 12
	b := "sunrisme"
	var c int

	var d *int

	fmt.Println("a 变量地址：", &a)
	fmt.Println("b 变量地址：", &b)
	fmt.Println("c 的地址：", &c)

	//指针变量
	var aPoint *int
	aPoint = &a
	fmt.Println("指针指向的内存地址中的值：", *aPoint)

	fmt.Println("指针指向的内存地址中的值: ", d)
}
