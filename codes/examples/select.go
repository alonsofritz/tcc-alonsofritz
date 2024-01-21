package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine produtora 1
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			ch1 <- fmt.Sprintf("Mensagem do Canal 1: %d", i)
		}
		close(ch1)
	}()

	// Goroutine produtora 2
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(2 * time.Second)
			ch2 <- fmt.Sprintf("Mensagem do Canal 2: %d", i)
		}
		close(ch2)
	}()

	// Goroutine consumidora com select
	for {
		select {
		case msg, ok := <-ch1:
			if ok {
				fmt.Println(msg)
			} else {
				fmt.Println("Canal 1 fechado.")
			}
		case msg, ok := <-ch2:
			if ok {
				fmt.Println(msg)
			} else {
				fmt.Println("Canal 2 fechado.")
			}
		}
	}
}
