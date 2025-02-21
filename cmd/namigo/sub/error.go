package sub

import "errors"

var (
	ErrMissingTerm      = errors.New("search term is missing")
	ErrPortfolioFailure = errors.New("portfolio collection failed")
)
