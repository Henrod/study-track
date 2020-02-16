package errors

import "errors"

var (
	ErrInternal              = errors.New("internal server error")
	ErrMissingImplementation = errors.New("missing implementation")
	ErrAlreadyExists         = errors.New("entity already exists")
	ErrNotFound              = errors.New("entity not found")
)

func New(msg string) error {
	return errors.New(msg)
}
