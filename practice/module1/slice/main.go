package main

import (
	"fmt"
)

// go语言中没有对应删除slice元素方法，可自行定义
func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...) // 切片后面...表示将切片打散作为参数传入函数
}

func main() {
	myArry := [5]int{1, 2, 3, 4, 5}
	mySlice := myArry[1:3]
	fmt.Printf("mySlice %+v\n", mySlice) // %+v 变量若为结构体会打印属性字段名
	fullSlice := myArry[:]
	remove3rdItem := deleteItem(fullSlice, 2)
	fmt.Printf("remove3rdItem %+v\n", remove3rdItem)
}
