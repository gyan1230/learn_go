package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
}

type square struct {
	side int
}

func (r rectangle) area() int {
	return r.length * r.breadth
}

func (s square) area() int {
	return s.side * s.side
}

type shape interface {
	area() int
}

func detail(s shape) {
	fmt.Println(s)
	fmt.Println("Area is ", s.area())
}

func main() {
	r := rectangle{
		length:  10,
		breadth: 5,
	}

	s := square{7}
	detail(r)
	detail(s)

}
