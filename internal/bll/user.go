package bll

import (
	"context"
)

type User struct {
	Name string
}

func CreateUser(
	ctx context.Context,
	storage Storage,
	user User,
) (res User, err error) {
	err = storage.CreateUser(ctx, user)
	if err != nil {
		return res, err
	}

	return user, nil
}

func GetUser(
	ctx context.Context,
	storage Storage,
	name string,
) (res User, err error) {
	user, err := storage.GetUser(ctx, name)
	if err != nil {
		return res, err
	}

	return user, nil
}
