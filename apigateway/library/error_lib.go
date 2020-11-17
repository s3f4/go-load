package library

import (
	"errors"
)

var (
	// ErrNotFound not found
	ErrNotFound error = errors.New("Error Not Found")
	// ErrUnauthorized login required
	ErrUnauthorized = errors.New("Unauthorized")
	// ErrUnprocessableEntity for bad requests
	ErrUnprocessableEntity = errors.New("Bad Request")
	// ErrBadRequest for bad requests
	ErrBadRequest = errors.New("Bad Request")
)
