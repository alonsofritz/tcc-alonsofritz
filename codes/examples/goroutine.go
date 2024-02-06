func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printMessage("Hello from Goroutine 1")
	}()

	go func() {
		defer wg.Done()
		printMessage("Hello from Goroutine 2")
	}()

	wg.Wait()
	fmt.Println("Main function exiting")
}

func printMessage(message string) {
	fmt.Println(message)
}
