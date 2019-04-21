package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func main() {
	wg.Add(2)
	go add("foo")
	go add("bar")
	wg.Wait()
	fmt.Println("Final counter: ", counter)
}

func add(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(s, i, "counter is ", counter)
		mutex.Unlock()
	}
	wg.Done()

}
