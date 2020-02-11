package errors

import "errors"

var (
	ErrInternal              = errors.New("internal server error")
	ErrMissingImplementation = errors.New("missing implementation")
)

func New(msg string) error {
	return errors.New(msg)
}
