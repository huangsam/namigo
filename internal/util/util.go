// Package util provides utility functions for the application.
package util

// dismiss ignores the error from a function/method
func dismiss(f func() error) {
	_ = f()
}
