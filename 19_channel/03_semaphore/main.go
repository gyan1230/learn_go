package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// go func() {
	// 	<-done
	// 	<-done
	// //	close(c)
	// }()

	// for n := range c {
	// 	fmt.Println(n)
	// }

	count := 0
outer:
	for {
		select {
		case n := <-c:
			fmt.Println(n)
		case <-done:
			count++
			if count == 2 {
				close(c)
				close(done)
				break outer
			}
		}
	}
}
