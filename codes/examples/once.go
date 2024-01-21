package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	// Função que será executada apenas uma vez
	initializer := func() {
		fmt.Println("Executando inicialização.")
	}

	// Goroutine 1
	go func() {
		once.Do(initializer)
		fmt.Println("Goroutine 1 concluída.")
	}()

	// Goroutine 2
	go func() {
		once.Do(initializer)
		fmt.Println("Goroutine 2 concluída.")
	}()

	// Goroutine 3
	go func() {
		once.Do(initializer)
		fmt.Println("Goroutine 3 concluída.")
	}()

	// Aguarda a conclusão das goroutines
	select {}
}
