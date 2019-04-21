package main

import "fmt"

func main() {

	n := fanin(display("gyan"), display("chand"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-n)
	}
	fmt.Println("Done...")
}

func display(name string) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s-%d", name, i)
		}
	}()
	return out
}

func fanin(in1, in2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			out <- <-in1
		}

	}()

	go func() {
		for {
			out <- <-in2
		}

	}()
	return out
}
