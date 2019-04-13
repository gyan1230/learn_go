package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	fmt.Println(m1)
	change(m1)
	fmt.Println(m1["gyan"])
}

func change(z map[string]int) {
	z["gyan"] = 12
}
