package main

import "fmt"

func main() {
	var a int
	a = 32

	fmt.Printf("%p - %d\n", &a, &a)
	fmt.Println(&a)
}
