package main

import (
	"runtime"
	"sync"
)

func main() {
	var mu sync.Mutex
	var items = make(map[int]struct{})

	runtime.SetMutexProfileFraction(5)
	var wg sync.WaitGroup
	for i := 0; i < 100_1000; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()

			items[i] = struct{}{}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
