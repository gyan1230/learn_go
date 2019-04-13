package main

import "fmt"

func main() {
	fmt.Println("This is about map -- key, value pair")
	m := map[string]int{
		"gyan":  12,
		"saanj": 30,
	}
	fmt.Println(m)
	fmt.Println(m["gyan"])

	// comma ok idiom
	v, ok := m["chand"]
	fmt.Println(v, ok)

	//to check whether a key is available
	if v, ok := m["gyan"]; ok {
		fmt.Println("check pass, key available and value is ", v)
	} else {
		fmt.Println("check failed, no key present...")
	}

	if v, ok := m["chand"]; ok {
		fmt.Println("check pass, key available and value is ", v)
	} else {
		fmt.Println("check failed, no key present...")
	}
}
