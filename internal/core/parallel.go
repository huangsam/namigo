package core

import "sync"

// goroutineCount is the common number of goroutines to start.
const goroutineCount = 10

// dynamicGoroutineCount returns goroutines based on workload size.
func dynamicGoroutineCount(workloadSize int) int {
	switch {
	case workloadSize <= 5:
		return 3
	case workloadSize <= 20:
		return 5
	case workloadSize <= 50:
		return 8
	default:
		return 10
	}
}

// StartCommonWorkers starts a common number of goroutines for a worker function.
func StartCommonWorkers(worker func()) {
	StartWorkers(goroutineCount, worker)
}

// StartDynamicWorkers starts goroutines based on workload size.
func StartDynamicWorkers(workloadSize int, worker func()) {
	count := dynamicGoroutineCount(workloadSize)
	StartWorkers(count, worker)
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
