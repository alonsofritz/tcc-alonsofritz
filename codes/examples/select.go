func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			ch1 <- fmt.Sprintf("Message from Channel 1: %d", i)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(2 * time.Second)
			ch2 <- fmt.Sprintf("Message from Channel 2: %d", i)
		}
		close(ch2)
	}()

	for {
		select {
		case msg, ok := <-ch1:
			if ok {
				fmt.Println(msg)
			} else {
				fmt.Println("Channel 1 closed.")
			}
		case msg, ok := <-ch2:
			if ok {
				fmt.Println(msg)
			} else {
				fmt.Println("Channel 2 closed.")
			}
		}
	}
}
