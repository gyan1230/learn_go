package main

import "fmt"

func main() {
	n := avg(1, 2, 3, 4, 5, 6)
	fmt.Println(n)
	data := []float64{1.2, 2.4, 3.6}
	n2 := avg2(data...)
	fmt.Println(n2)

}

func avg(input ...float64) float64 {
	total := 0.0
	for _, v := range input {
		total += v
	}
	return total / float64(len(input))
}

func avg2(in2 ...float64) float64 {
	var total float64
	for _, v := range in2 {
		total += v
	}
	return total / float64(len(in2))
}
