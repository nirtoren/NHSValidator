package validator

import "errors"

var (
	ErrBadInput = errors.New("wrong Input, must be 10 digits")
	ErrDivisionNotAllowed = errors.New("error on dividing operation")
	ErrInvalidNHS = errors.New("invalid NHS number")
)
