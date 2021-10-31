package main

import (
	"flag"
	"fmt"
)

func main() {
	age := flag.Int("age", 18, "age")
	sex := flag.String("sex", "male", "sex")
	flag.Parse()

	switch *sex {
	case "male":
		if *age > 18 {
			fmt.Println("man")
		} else {
			fmt.Println("boy")
		}

	case "female":
		if *age > 18 {
			fmt.Println("women")
		} else {
			fmt.Println("girl")
		}

	}
}
