package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type detail struct {
	person
	company string
}

func main() {

	emp1 := detail{
		person: person{
			first: "gyan",
			last:  "mohanty",
			age:   31,
		},
		company: "raddef",
	}
	fmt.Println(emp1.first, emp1.last, emp1.age, emp1.company)
}
