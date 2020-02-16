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

func GetUser(
	ctx context.Context,
	logger logrus.FieldLogger,
	storage Storage,
	name string,
) (res User, err error) {
	user, err := storage.GetUser(ctx, name)
	if err != nil {
		logger.WithError(err).Errorf("failed to get user")
		return res, fmt.Errorf("failed to create user: %w", errors.ErrInternal)
	}

	return user, nil
}
