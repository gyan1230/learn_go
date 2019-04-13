package main

import "fmt"

func main() {
	x := make([]int, 5, 10)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))
	for i := 0; i < 1024*1024; i++ {
		x = append(x, i)
		//	fmt.Println(x)
		if len(x) == cap(x) {
			fmt.Println(len(x), cap(x))

		}

	}
}
