package search

import "errors"

var (
	ErrPorftolioEmpty   = errors.New("portfolio collection empty")
	ErrPorftolioFailure = errors.New("portfolio collection failure")
)
