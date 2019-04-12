package main

import (
	"fmt"
)

func main() {
	switch 1 + 2 {
	case 1:
		fmt.Println("case : 1")
		//		fallthrough
	case 2:
		fmt.Println("case : 2")
		//		fallthrough
	case 3:
		fmt.Println("case : 3")
		//		fallthrough
	default:
		fmt.Println("Default statements...")
	}
}
