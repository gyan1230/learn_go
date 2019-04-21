package main

import "fmt"

const (
	_ = iota

	kB = 1 << (iota * 10)
	mB = 1 << (iota * 10)
	gB = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%b - KB is :%d bytes\n", kB, kB)
	fmt.Printf("%b - MB is :%d bytes\n", mB, mB)
	fmt.Printf("%b - GB is :%d bytes\n", gB, gB)
}
