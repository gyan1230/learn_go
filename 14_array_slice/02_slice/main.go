package main

import "fmt"

func main() {
	var x []int
	y := []int{12, 23, 45, 56}
	fmt.Println("Appending slice ...")
	fmt.Println(x[:])
	x = append(x, 1, 2, 3)
	fmt.Println(x)
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("********************")
	fmt.Println(x[:])
	fmt.Println(x[:8])
	fmt.Println(y[:4])
	fmt.Println(x[3:4])
	fmt.Println(y[1:3])

	fmt.Println("Deleting 2, 3 from slice X...")
	x = append(x[:1], x[3:]...)
	fmt.Println(x)

}
