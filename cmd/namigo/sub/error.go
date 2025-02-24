package sub

import "errors"

var (
	// User does not provide any search term
	ErrMissingSearchTerm = errors.New("missing search term")

	// No results were collected
	ErrPorftolioEmpty = errors.New("portfolio collection empty")

	// At least one failure was identified
	ErrPorftolioFailure = errors.New("portfolio collection failure")
)
