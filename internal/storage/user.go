package storage

import (
	"time"

	"github.com/Henrod/study-track/internal/bll"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromBLLUser(user bll.User) User {
	return User{
		ID:   user.ID,
		Name: user.Name,
	}
}
