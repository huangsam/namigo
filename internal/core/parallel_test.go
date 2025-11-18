package core

import (
	"sync/atomic"
	"testing"
)

func TestStartWorkers(t *testing.T) {
	var counter int64
	StartWorkers(5, func() {
		atomic.AddInt64(&counter, 1)
	})
	if counter != 5 {
		t.Errorf("expected 5, got %d", counter)
	}
}

func TestStartCommonWorkers(t *testing.T) {
	var counter int64
	StartCommonWorkers(func() {
		atomic.AddInt64(&counter, 1)
	})
	if counter != 10 {
		t.Errorf("expected 10, got %d", counter)
	}
}
