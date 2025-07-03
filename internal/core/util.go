// Package core provides utility functions for the application.
package core

// dismiss ignores the error from a function/method
func dismiss(f func() error) {
	_ = f()
}
