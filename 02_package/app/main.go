package main

import (
	"fmt"

	"github.com/gyan1230/learn_go/02_package/gyan"
)

func main() {
	fmt.Println("Package creation...")
	fmt.Println(gyan.A)
	fmt.Println(gyan.Deduct(20))
}
