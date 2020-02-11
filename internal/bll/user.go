package bll

import (
	"context"
	"fmt"

	"github.com/Henrod/study-track/internal/storage/db"

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
	storage db.Querier,
	user User,
) (res User, err error) {
	created, err := storage.CreateUser(ctx, db.CreateUserParams{
		ID:   uuid.New(),
		Name: user.Name,
	})
	if err != nil {
		logger.WithError(err).Errorf("failed to create user")
		return res, fmt.Errorf("failed to create user: %w", errors.ErrInternal)
	}

	return User{
		ID:   created.ID,
		Name: created.Name,
	}, nil
}
