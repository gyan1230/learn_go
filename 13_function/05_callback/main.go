package main

import "fmt"

func visit(in []int, callback func(int)) {
	for _, v := range in {
		callback(v)

	}
}

func main() {
	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})
}
