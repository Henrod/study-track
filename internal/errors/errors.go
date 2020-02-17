package errors

import (
	"errors"
)

var (
	ErrInternal              = errors.New("internal server error")
	ErrMissingImplementation = errors.New("missing implementation")
	ErrAlreadyExists         = errors.New("entity already exists")
	ErrNotFound              = errors.New("entity not found")
	ErrInvalidArgument       = errors.New("invalid argument")
)

func New(msg string) error {
	return errors.New(msg)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
