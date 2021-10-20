package main

import (
	"flag" //用于解析传入参数
	"fmt"
	"os"
)

func main() {
	// 解析命令行传参 --name 默认值为“world”, 描述字符串 "specify the name you want to sya hi"
	name := flag.String("name", "world", "specify the name you want to sya hi")
	flag.Parse()
	fmt.Println("os args is", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)
}
