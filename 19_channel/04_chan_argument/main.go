package main

import (
	"fmt"
	"time"
)

func main() {
	c := increment()
	p := puller(c)
	time.Sleep(time.Second)
	fmt.Println("before final value....")

	for n := range p {
		time.Sleep(time.Second)
		fmt.Println("from puller function:", n)
	}

}

func increment() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
			// fmt.Println("increment out.", i)
		}
		time.Sleep(time.Second)
		fmt.Println("close channel inside increment go routine...")
		close(out)

	}()
	time.Sleep(time.Second)
	fmt.Println("increment return ", out)
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)

	go func() {
		var sum int
		for n := range c {
			sum += n
			// fmt.Println("puller sum.", sum)
		}
		time.Sleep(time.Second)
		fmt.Println("close channel inside puller go routine...")
		out <- sum
		close(out)

	}()
	time.Sleep(time.Second)
	fmt.Println("puller return ", out)
	return out
}
