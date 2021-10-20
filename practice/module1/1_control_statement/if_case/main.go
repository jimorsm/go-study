package main

import (
	"fmt"
	"flag"
)


func main() {
	a := 123
	if a > 2 {
		fmt.Println(a)
	}
	switch a {
	case 10:
		fmt.Println(a)

	case 100:                    
		fmt.Println(a)

	default:
		fmt.Println(a)
	}
}
