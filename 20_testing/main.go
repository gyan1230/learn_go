package main

import "fmt"

func main() {
	fmt.Println("1+2 :", mysum(1, 2))
	fmt.Println("1+12 :", mysum(1, 12))
	fmt.Println("13+4+2 :", mysum(13, 4, 2))
}

func mysum(in1 ...int) int {
	var sum int
	for _, v := range in1 {
		sum += v
	}
	return sum
}
