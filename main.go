package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")

			// do some work
			fmt.Println(i, "working")

			m.Unlock()
			fmt.Println(i, "done")
			wg.Done()
		}(i)
	}

	wg.Wait()
}
