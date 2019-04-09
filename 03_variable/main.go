package main

import "fmt"

func main() {
	shortHand()
	zeroValue()
}

func shortHand() {
	a := 1
	b := 2.3
	c := "gyan"
	d := true
	fmt.Printf("Short hand declaration...\n")

	fmt.Printf("%v = %T\n", a, a)
	fmt.Printf("%v = %T\n", b, b)
	fmt.Printf("%v = %T\n", c, c)
	fmt.Printf("%v = %T\n", d, d)
}

func zeroValue() {
	var aa int
	var bb float32
	var cc string
	var dd bool

	fmt.Printf("Zero Value declaration...\n")

	fmt.Printf("%v = %T\n", aa, aa)
	fmt.Printf("%v = %T\n", bb, bb)
	fmt.Printf("%v = %T\n", cc, cc)
	fmt.Printf("%v = %T\n", dd, dd)

}
