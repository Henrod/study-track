package server

import "context"

type Storage interface {
	CreateStudent(ctx context.Context, student Student) (err error)
	CreateTheme(ctx context.Context, theme Theme) (err error)
	CreateSubject(ctx context.Context, subject Subject) (err error)

	GetStudent(ctx context.Context, name string) (student Student, err error)
	ListThemes(_ context.Context, name string) (res []Theme, err error)
}
