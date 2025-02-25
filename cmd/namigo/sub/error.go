package sub

import "errors"

var (
	ErrMissingSearchTerm = errors.New("missing search term")
	ErrPorftolioEmpty    = errors.New("portfolio collection empty")
	ErrPorftolioFailure  = errors.New("portfolio collection failure")
)
