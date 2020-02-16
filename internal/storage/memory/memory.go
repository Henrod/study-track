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
	users     = map[uuid.UUID]storage.User{}
	usersName = map[string]uuid.UUID{}
	subjects  = map[uuid.UUID]storage.Subject{}
)

type Storage struct{}

var _ bll.Storage = &Storage{}

func (s *Storage) CreateUser(_ context.Context, user bll.User) (err error) {
	user.ID = uuid.New()

	if _, exists := usersName[user.Name]; exists {
		return fmt.Errorf("user already exists: %w", errors.ErrAlreadyExists)
	}

	users[user.ID] = storage.FromBLLUser(user)
	usersName[user.Name] = user.ID

	return nil
}

func (s *Storage) GetUser(_ context.Context, name string) (user bll.User, err error) {
	userID, exists := usersName[name]
	if !exists {
		return user, fmt.Errorf("user not found: %w", errors.ErrNotFound)
	}

	return users[userID].ToBLLUser(), nil
}

func (s *Storage) CreateSubject(_ context.Context, sub bll.Subject) (err error) {
	sub.ID = uuid.New()
	subjects[sub.ID] = storage.FromBLLSubject(sub)
	return nil
}
