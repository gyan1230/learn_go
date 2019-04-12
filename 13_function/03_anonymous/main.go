package main

import "fmt"

func main() {
	welcome := func() {
		fmt.Println("welcome go ...")
	}
	welcome()
	fmt.Printf("type is %T\n", welcome)

	//error in return from anonymous function

	rcv := func() string {
		s := "gyan"
		return s
	}

	fmt.Println(rcv())

	increment := func(x int) int {
		x++
		return x
	}
	fmt.Println(increment(5))
	fmt.Println(increment(25))

}
