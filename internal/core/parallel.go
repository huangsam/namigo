package core

import "sync"

// goroutineCount is the common number of goroutines to start.
const goroutineCount = 10

// StartCommonWorkers starts a common number of goroutines for a worker function.
func StartCommonWorkers(worker func()) {
	StartWorkers(goroutineCount, worker)
}

// StartWorkers starts a number of goroutines for a worker function.
func StartWorkers(count int, worker func()) {
	var wg sync.WaitGroup
	for range count {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker()
		}()
	}
	wg.Wait()
}
