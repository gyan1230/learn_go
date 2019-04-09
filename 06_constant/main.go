package main

import "fmt"

const (
	_ = iota

	kB = 1 << (iota * 10)
	mB = 1 << (iota * 10)
	gB = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%b - %d\n", kB, kB)
	fmt.Printf("%b - %d\n", mB, mB)
	fmt.Printf("%b - %d\n", gB, gB)
}
