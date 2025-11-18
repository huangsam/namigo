package core

import (
	"errors"
	"testing"
)

func TestDismiss(t *testing.T) {
	called := false
	f := func() error {
		called = true
		return errors.New("test error")
	}
	dismiss(f)
	if !called {
		t.Error("function was not called")
	}
}
