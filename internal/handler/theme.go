package handler

import (
	"context"

	"github.com/Henrod/study-track/internal/bll"
	"github.com/Henrod/study-track/pkg/studytrack"
	"github.com/sirupsen/logrus"
)

type Theme struct {
	storage bll.Storage
	logger  logrus.FieldLogger
}

var _ studytrack.ThemeServiceServer = &Theme{}

func NewTheme(storage bll.Storage, logger logrus.FieldLogger) *Theme {
	return &Theme{storage: storage, logger: logger}
}

func (u *Theme) Create(
	ctx context.Context,
	req *studytrack.CreateThemeRequest,
) (res *studytrack.Theme, err error) {
	_, err = bll.CreateTheme(ctx, u.storage, bll.Theme{
		UserName: req.User,
		Name:     req.Theme.Name,
	})
	if err != nil {
		return res, parseErr(err, "failed to create theme")
	}

	res = &studytrack.Theme{
		Name: req.Theme.Name,
		User: &studytrack.User{
			Name: req.User,
		},
	}

	return res, nil
}

func (u *Theme) List(
	ctx context.Context,
	req *studytrack.ListThemesRequest,
) (res *studytrack.ListThemesResponse, err error) {
	themes, err := bll.ListThemes(ctx, u.storage, req.User)
	if err != nil {
		return res, parseErr(err, "failed to list themes")
	}

	res = &studytrack.ListThemesResponse{
		Themes: []*studytrack.Theme{},
	}

	for _, t := range themes {
		res.Themes = append(res.Themes, &studytrack.Theme{
			Name: t.Name,
			User: &studytrack.User{
				Name: t.UserName,
			},
		})
	}

	return res, nil
}
