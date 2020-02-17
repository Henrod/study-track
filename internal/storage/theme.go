package storage

import (
	"github.com/Henrod/study-track/internal/bll"
	"github.com/google/uuid"
)

type Theme struct {
	ID       uuid.UUID
	UserName string
	Name     string
}

func (u Theme) ToBLLTheme() bll.Theme {
	return bll.Theme{
		ID:       u.ID,
		UserName: u.UserName,
		Name:     u.Name,
	}
}

func FromBLLTheme(theme bll.Theme) Theme {
	return Theme{
		ID:       theme.ID,
		UserName: theme.UserName,
		Name:     theme.Name,
	}
}
