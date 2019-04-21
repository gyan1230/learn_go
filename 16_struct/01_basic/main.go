package main

import "fmt"

func main() {
	type person struct {
		first string
		last  string
		age   int
	}

	gm := person{
		first: "gyan chand",
		last:  "mohanty",
		age:   31,
	}

	ak := person{
		first: "ajay",
		last:  "kumar",
		age:   29,
	}

	ps := person{
		first: "parath",
		last:  "singh",
		age:   28,
	}

	fmt.Println(gm, ak, ps)
	fmt.Println(gm.first, gm.last, gm.age)
	fmt.Println(ak.first, ak.last, ak.age)
	fmt.Println(ps.first, ps.last, ps.age)
}
