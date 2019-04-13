package main

import "fmt"

func main() {
	gm := []string{"gyan", "mohanty", "bike", "girls", "liquor"}
	ps := []string{"Parath", "Singh", "aunty", "weed", "food"}

	xp := [][]string{gm, ps}
	fmt.Println(xp)
	fmt.Println(xp[1][1])
}
