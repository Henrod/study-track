package memory

import (
	"context"
	"fmt"

	"github.com/Henrod/study-track/internal/errors"

	"github.com/Henrod/study-track/internal/bll"
	"github.com/Henrod/study-track/internal/storage"
	"github.com/google/uuid"
)

var (
	users    = map[string]storage.User{}
	subjects = map[uuid.UUID]storage.Subject{}
	themes   = map[uuid.UUID]storage.Theme{}
)

type Storage struct{}

var _ bll.Storage = &Storage{}

func (s *Storage) CreateUser(_ context.Context, user bll.User) (err error) {
	if user.Name == "" {
		return fmt.Errorf("empty user name: %w", errors.ErrInvalidArgument)
	}

	if _, exists := users[user.Name]; exists {
		return fmt.Errorf("user already exists: %w", errors.ErrAlreadyExists)
	}

	users[user.Name] = storage.FromBLLUser(user)

	return nil
}

func (s *Storage) CreateSubject(_ context.Context, sub bll.Subject) (err error) {
	sub.ID = uuid.New()
	subjects[sub.ID] = storage.FromBLLSubject(sub)
	return nil
}

func (s *Storage) CreateTheme(_ context.Context, theme bll.Theme) (err error) {
	theme.ID = uuid.New()
	themes[theme.ID] = storage.FromBLLTheme(theme)
	return nil
}

func (s *Storage) GetUser(_ context.Context, name string) (user bll.User, err error) {
	u, exists := users[name]
	if !exists {
		return user, fmt.Errorf("user not found: %w", errors.ErrNotFound)
	}

	return u.ToBLLUser(), nil
}

func (s *Storage) ListThemes(_ context.Context, user string) (res []bll.Theme, err error) {
	for _, theme := range themes {
		if theme.UserName == user {
			res = append(res, theme.ToBLLTheme())
		}
	}

	return res, nil
}
