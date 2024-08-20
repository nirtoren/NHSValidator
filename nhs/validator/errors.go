package validator

import "errors"

var (
	ErrBadInput = errors.New("wrong Input, must be 10 digits")
	ErrDivisionNotAllowed = errors.New("error on dividing operation")
	// ErrInitialization = errors.New("error while initializing component")
	ErrInvalidNHS = errors.New("invalid NHS number")
)
