package main

import "fmt"

//package level
var x = 32
var y = "gyan"

func main() {
	c := 3.2
	fmt.Println(x, y, c)
	//block level scope
	var a = 10
	display()
	fmt.Println(a)

}

func display() {
	fmt.Println(x, y)
}
