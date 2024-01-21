package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)
	wg.Add(2)

	go incrementCounter()
	go incrementCounter()

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

func incrementCounter() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		// Critic! Race condition here!
		counter++
	}
}
