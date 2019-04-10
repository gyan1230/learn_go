package main

import "fmt"

var a int

type gyan int

func main() {
	a = 10
	var b gyan
	b = 20
	fmt.Println("value of a is ", a, "\nvalue of b is ", b)
	fmt.Printf("a type is %T\n", a)
	fmt.Printf("b type is %T\n", b)
	a = int(b)

	fmt.Println("value of a is ", a, "\nvalue of b is ", b)

}
