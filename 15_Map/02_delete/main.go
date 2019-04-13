package main

import "fmt"

func main() {
	m := map[string]int{
		"gyan":   1230,
		"ajay":   4567,
		"parath": 2611,
	}
	fmt.Println(m)
	m["sri"] = 1987
	for k, v := range m {
		fmt.Println(k, "-", v)
	}
	fmt.Println("*****************")
	fmt.Println("Sri deleted....")
	delete(m, "sri")
	for k, v := range m {
		fmt.Println(k, "-", v)
	}
}
