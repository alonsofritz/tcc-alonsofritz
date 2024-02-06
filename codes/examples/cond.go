type Data struct {
	value int
	cond  *sync.Cond
}

func main() {
	data := &Data{
		value: 0,
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	go func() {
		for {
			time.Sleep(2 * time.Second)
			data.cond.L.Lock()
			data.value++
			fmt.Println("Updated value:", data.value)
			data.cond.Signal()
			data.cond.L.Unlock()
		}
	}()

	go func() {
		for {
			data.cond.L.Lock()
			data.cond.Wait()
			fmt.Println("Updated value:", data.value)
			data.cond.L.Unlock()
		}
	}()

	select {}
}
