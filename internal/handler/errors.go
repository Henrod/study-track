package handler

import (
	"github.com/Henrod/study-track/internal/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func parseErr(err error, msg string) (res error) {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, errors.ErrNotFound):
		return status.Errorf(codes.NotFound, "%s: %v", msg, err)

	case errors.Is(err, errors.ErrAlreadyExists):
		return status.Errorf(codes.AlreadyExists, "%s: %v", msg, err)

	case errors.Is(err, errors.ErrMissingImplementation):
		return status.Errorf(codes.Internal, "%s: %v", msg, err)
	}

	return errors.New(msg)
}
