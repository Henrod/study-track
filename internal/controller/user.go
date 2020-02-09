package controller

import (
	"context"

	"github.com/Henrod/study-track/internal/bll"
	"github.com/Henrod/study-track/pkg/studytrack"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type User struct {
	storage bll.Storage
	logger  logrus.FieldLogger
}

func NewUser(storage bll.Storage, logger logrus.FieldLogger) *User {
	return &User{storage: storage, logger: logger}
}

func (u *User) Create(
	ctx context.Context,
	req *studytrack.CreateUserRequest,
) (res *studytrack.User, err error) {
	user := bll.User{
		ID:   uuid.UUID{},
		Name: req.User.Name,
	}

	err = bll.CreateUser(ctx, u.logger, u.storage, user)
	if err != nil {
		return res, err
	}

	res = &studytrack.User{
		Name: user.Name,
	}

	return res, nil
}
