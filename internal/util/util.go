package util

// dismiss ignores the error from a function/method
func dismiss(f func() error) {
	_ = f()
}
