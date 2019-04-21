package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "gyan chand",
		last:  "mohanty",
		age:   31,
	}

	fmt.Println(p1.first, p1.last, p1.age)
}
