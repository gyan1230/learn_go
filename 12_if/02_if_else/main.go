package main

import "fmt"

func main() {
	if b := "red"; !true {
		fmt.Println(b)
	} else {
		fmt.Println("false")
	}
	//can't access b here scope issues
	// fmt.Println(b)
}
