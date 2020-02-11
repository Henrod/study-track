// This package is only supposed to test bll

package memory

import (
	"context"

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
	user.ID = uuid.New()
	users[user.ID] = storage.FromBLLUser(user)
	return nil
}

func (s *Storage) CreateSubject(_ context.Context, sub bll.Subject) (err error) {
	sub.ID = uuid.New()
	subjects[sub.ID] = storage.FromBLLSubject(sub)
	return nil
}
