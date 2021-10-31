package main

import (
	"fmt"
)

func main() {
	mySlice1 := new([]int)     // new 根据传入类型分配一片内存空间，并返回指向这片内存空间的指针
	mySlice2 := make([]int, 0) // 初始化内置的数据结构比如切片、哈希表和 Channel
	mySlice3 := make([]int, 10)
	mySlice4 := make([]int, 10, 20)
	fmt.Printf("mySlice1 %v\n", mySlice1)
	fmt.Printf("mySlice2 %v\n", mySlice2)
	fmt.Printf("mySlice3 %v\n", mySlice3)
	fmt.Printf("mySlice4 %v\n", mySlice4)
}
