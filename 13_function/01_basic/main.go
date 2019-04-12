package main

import "fmt"

func main() {
	welcome("Gyan")
	fmt.Println(welcome2("Gyanchand ", "mohanty"))
	fmt.Println(welcome3("Dipu", "Saanj"))
}

func welcome(name string) {
	fmt.Println("welcome ", name)
}

func welcome2(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func welcome3(s1, s2 string) (string, string) {
	return fmt.Sprint(s1, s2), fmt.Sprint(s2, s1)
}
