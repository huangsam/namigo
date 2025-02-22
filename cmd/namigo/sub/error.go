package sub

import "errors"

var (
	ErrMissingTerm    = errors.New("search term is missing")
	ErrPorftolioEmpty = errors.New("portfolio collection is empty")
)
