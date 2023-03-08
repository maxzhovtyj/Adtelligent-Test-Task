package models

import "errors"

var (
	ErrInvalidAuthorizationHeader = errors.New("invalid Authorization header")
	ErrInvalidInputBody           = errors.New("invalid input body")
)
