func main() {
	messageChannel := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			message := fmt.Sprintf("Mensagem %d", i)
			messageChannel <- message
		}
		close(messageChannel)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for message := range messageChannel {
			fmt.Println("Recebido:", message)
		}
	}()

	wg.Wait()
}
