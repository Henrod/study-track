package errors

import "errors"

var (
	ErrInternal = errors.New("internal server error")
)

func New(msg string) error {
	return errors.New(msg)
}
