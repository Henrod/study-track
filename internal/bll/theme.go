package bll

import (
	"context"

	"github.com/google/uuid"
)

type Theme struct {
	ID       uuid.UUID
	UserName string
	Name     string
}

func CreateTheme(
	ctx context.Context,
	storage Storage,
	theme Theme,
) (res Theme, err error) {
	err = storage.CreateTheme(ctx, theme)
	if err != nil {
		return res, err
	}

	return theme, nil
}

func ListThemes(
	ctx context.Context,
	storage Storage,
	user string,
) (res []Theme, err error) {
	themes, err := storage.ListThemes(ctx, user)
	if err != nil {
		return res, err
	}

	return themes, nil
}
