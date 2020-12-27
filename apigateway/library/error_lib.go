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
	// ErrRefreshTokenExpire refresh token expired
	ErrRefreshTokenExpire = errors.New("Refresh token is expired")
	// ErrInternalServerError internal errors
	ErrInternalServerError = errors.New("Internal server error")
	// ErrForbidden forbidden error
	ErrForbidden = errors.New("Forbidden")
)
