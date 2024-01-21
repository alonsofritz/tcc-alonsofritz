package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	value int
	cond  *sync.Cond
}

func main() {
	data := &Data{
		value: 0,
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	// Goroutine para atualizar o valor após um certo tempo
	go func() {
		for {
			time.Sleep(2 * time.Second)
			data.cond.L.Lock()
			data.value++
			fmt.Println("Valor atualizado:", data.value)
			data.cond.Signal() // Sinaliza que o valor foi atualizado
			data.cond.L.Unlock()
		}
	}()

	// Goroutine para imprimir o valor quando atualizado
	go func() {
		for {
			data.cond.L.Lock()
			data.cond.Wait() // Aguarda o sinal de atualização
			fmt.Println("Valor atual:", data.value)
			data.cond.L.Unlock()
		}
	}()

	// Aguarda indefinidamente
	select {}
}
