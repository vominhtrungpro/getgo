package inventory

import "errors"

var (
	// ErrNotFound is error when no item found in db
	ErrNotFound = errors.New("not found")
)
