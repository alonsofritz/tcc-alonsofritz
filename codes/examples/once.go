func main() {
	var once sync.Once

	initializer := func() {
		fmt.Println("Initializing...")
	}

	go func() {
		once.Do(initializer)
		fmt.Println("Hello from Goroutine 1.")
	}()

	go func() {
		once.Do(initializer)
		fmt.Println("Hello from Goroutine 2.")
	}()

	go func() {
		once.Do(initializer)
		fmt.Println("Hello from Goroutine 3.")
	}()

	select {}
}
