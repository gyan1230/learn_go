package main

import "fmt"

// main function exit before other 2 function, that's why no output

func main() {
	go foo()
	go bar()
}
func foo() {
	for i := 0; i < 50; i++ {
		fmt.Println("foo: ", i)
	}
}

func bar() {
	for i := 0; i < 50; i++ {
		fmt.Println("bar: ", i)
	}
}
