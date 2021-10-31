package main

import (
	"fmt"
)

func main() {
	mySlice := []int{10, 20, 30, 40, 50}
	for _, value := range mySlice {
		value *= 2 // go语言中都是值传递，此处value为临时变量
	}
	fmt.Printf("mySlice %+v\n", mySlice)
	for index := range mySlice {
		mySlice[index] *= 2 // 使用index来操作切片元素
	}
	fmt.Printf("mySlice %+v\n", mySlice)
}
