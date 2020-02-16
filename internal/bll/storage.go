package bll

import "context"

type Storage interface {
	CreateUser(ctx context.Context, user User) (err error)
	GetUser(ctx context.Context, name string) (user User, err error)

	CreateSubject(ctx context.Context, subject Subject) (err error)
}
