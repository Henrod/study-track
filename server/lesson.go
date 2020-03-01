package server

import (
	"context"
	"strings"

	"github.com/Henrod/study-track/studytrack"
	"github.com/sirupsen/logrus"
)

type Lesson struct {
	storage Storage
	logger  logrus.FieldLogger
}

var _ studytrack.LessonsServer = &Lesson{}

func NewLesson(storage Storage, logger logrus.FieldLogger) *Lesson {
	return &Lesson{storage: storage, logger: logger}
}

func (l *Lesson) CreateTheme(
	ctx context.Context,
	req *studytrack.CreateThemeRequest,
) (res *studytrack.Theme, err error) {
	studentName := strings.Split(req.Parent, "/")[1]
	student, err := l.storage.GetStudent(ctx, studentName)
	if err != nil {
		return nil, parseErr(err, "failed to get parent student")
	}

	theme := (*Theme)(nil).fromProto(req.Theme, student.Name)
	err = l.storage.CreateTheme(ctx, *theme)
	if err != nil {
		return nil, parseErr(err, "failed to create theme")
	}

	return theme.toProto(), nil
}

func (l *Lesson) ListThemes(
	ctx context.Context,
	req *studytrack.ListThemesRequest,
) (res *studytrack.ListThemesResponse, err error) {
	studentName := strings.Split(req.Name, "/")[1]
	_, err = l.storage.GetStudent(ctx, studentName)
	if err != nil {
		return nil, parseErr(err, "failed to get student")
	}

	themes, err := l.storage.ListThemes(ctx, studentName)
	if err != nil {
		return nil, parseErr(err, "failed to list themes")
	}

	return &studytrack.ListThemesResponse{
		Themes: Themes(themes).toProto(),
	}, nil
}

func (l *Lesson) CreateSubject(context.Context, *studytrack.CreateSubjectRequest) (*studytrack.Subject, error) {
	panic("implement me")
}
