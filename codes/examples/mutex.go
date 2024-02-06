var count int
var mutex sync.Mutex

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	count++
	fmt.Printf("Incrementing: %d\n", count)
}

func decrement() {
	mutex.Lock()
	defer mutex.Unlock()
	count--
	fmt.Printf("Decrementing: %d\n", count)
}

func main() {
	var arithmetic sync.WaitGroup

	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	for i := 0; i < 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Final Count:", count)
}
