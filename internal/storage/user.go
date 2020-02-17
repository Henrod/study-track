package storage

import (
	"time"

	"github.com/Henrod/study-track/internal/bll"
)

type User struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) ToBLLUser() bll.User {
	return bll.User{
		Name: u.Name,
	}
}

func FromBLLUser(user bll.User) User {
	return User{
		Name: user.Name,
	}
}
