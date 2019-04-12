package main

import "fmt"

func main() {
	switch "gyan" {
	case "gyan", "chand":
		fmt.Println("case gyan or chand")
		fallthrough
	case "dipu", "saanj":
		fmt.Println("case dipu or saanj")
	case "suzuki", "yamaha":
		fmt.Println("case suzuki or yamaha")

	default:
		fmt.Println("Default case")
	}

}
