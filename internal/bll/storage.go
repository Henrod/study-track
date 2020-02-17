package bll

import "context"

type Storage interface {
	CreateUser(ctx context.Context, user User) (err error)
	CreateTheme(ctx context.Context, theme Theme) (err error)
	CreateSubject(ctx context.Context, subject Subject) (err error)

	GetUser(ctx context.Context, name string) (user User, err error)
	ListThemes(_ context.Context, user string) (res []Theme, err error)
}
