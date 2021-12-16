package err

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrCategoryNotFound = errors.New("category not found")
)
