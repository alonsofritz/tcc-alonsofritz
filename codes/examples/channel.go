package main

import (
	"fmt"
	"sync"
)

func main() {
	// Criando um canal de comunicação entre goroutines
	messageChannel := make(chan string)

	// WaitGroup para garantir que todas as goroutines sejam concluídas antes do término do programa
	var wg sync.WaitGroup

	// Goroutine produtora
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			message := fmt.Sprintf("Mensagem %d", i)
			messageChannel <- message // Enviando mensagem para o canal
		}
		close(messageChannel) // Fechando o canal após a conclusão do envio de mensagens
	}()

	// Goroutine consumidora
	wg.Add(1)
	go func() {
		defer wg.Done()
		for message := range messageChannel {
			fmt.Println("Recebido:", message) // Recebendo mensagem do canal
		}
	}()

	// Aguardando a conclusão de ambas as goroutines
	wg.Wait()
}
