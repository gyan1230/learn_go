package main

import "fmt"

func main() {
	// n is an array
	n := [5]int{1, 2, 3, 4, 5}
	fmt.Println(n)

	//m is a slice
	m := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	for i, v := range m {
		fmt.Println(i, "-", v)
	}
}
