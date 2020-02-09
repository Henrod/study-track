package memory

import (
	"context"

	"github.com/Henrod/study-track/internal/errors"

	"github.com/Henrod/study-track/internal/bll"
	"github.com/Henrod/study-track/internal/storage"
	"github.com/google/uuid"
)

var (
	users    = map[uuid.UUID]storage.User{}
	subjects = map[uuid.UUID]storage.Subject{}
)

type Storage struct{}

var _ bll.Storage = &Storage{}

func (s *Storage) CreateUser(_ context.Context, user bll.User) (err error) {
	if user.ID == uuid.Nil {
		return errors.New("empty user id")
	}

	users[user.ID] = storage.FromBLLUser(user)
	return nil
}

func (s *Storage) CreateSubject(_ context.Context, sub bll.Subject) (err error) {
	if sub.ID == uuid.Nil {
		return errors.New("empty user id")
	}

	subjects[sub.ID] = storage.FromBLLSubject(sub)
	return nil
}
