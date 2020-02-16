package bll

import (
	"context"
	"fmt"

	"github.com/Henrod/study-track/internal/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID   uuid.UUID
	Name string
}

func CreateUser(
	ctx context.Context,
	logger logrus.FieldLogger,
	storage Storage,
	user User,
) (res User, err error) {
	err = storage.CreateUser(ctx, user)
	if err != nil {
		logger.WithError(err).Errorf("failed to create user")
		return res, fmt.Errorf("failed to create user: %w", errors.ErrInternal)
	}

	return user, nil
}
