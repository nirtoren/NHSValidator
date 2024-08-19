package validator

import "errors"

var (
	ErrBadInput = errors.New("Wrong Input, must be 10 digits")
	ErrDivisionNotAllowed = errors.New("Error on dividing operation")
	ErrInitialization = errors.New("Error while initializing component")
	ErrInvalidNHS = errors.New("Invalid NHS number")
)
