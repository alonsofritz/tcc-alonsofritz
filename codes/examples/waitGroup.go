package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Adiciona 2 ao WaitGroup para indicar que estamos esperando duas goroutines
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done() // Decrementa o contador quando a goroutine termina
		printMessage("Hello from Goroutine 1")
	}()

	// Goroutine 2
	go func() {
		defer wg.Done() // Decrementa o contador quando a goroutine termina
		printMessage("Hello from Goroutine 2")
	}()

	// Espera até que todas as goroutines tenham terminado
	wg.Wait()

	fmt.Println("Main function exiting")
}

func printMessage(message string) {
	fmt.Println(message)
}
