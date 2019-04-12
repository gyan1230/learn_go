package main

import "fmt"

func main() {
	welcome := func() {
		fmt.Println("welcome go ...")
	}
	welcome()
	fmt.Printf("type is %T\n", welcome)

	//error in return from anonymous function
	/*
		rcv := func2() string{
			s:= "gyan"
			return s
		}

		fmt.Println(rcv())
	*/
}
