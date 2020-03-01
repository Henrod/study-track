package memory

import (
	"context"
	"fmt"

	"github.com/Henrod/study-track/server"

	"github.com/Henrod/study-track/errors"

	"github.com/google/uuid"
)

var (
	students = map[string]server.Student{}
	subjects = map[uuid.UUID]server.Subject{}
	themes   = map[uuid.UUID]server.Theme{}
)

type Storage struct{}

var _ server.Storage = &Storage{}

func (s *Storage) CreateStudent(_ context.Context, student server.Student) (err error) {
	if student.Name == "" {
		return fmt.Errorf("empty student name: %w", errors.ErrInvalidArgument)
	}

	if _, exists := students[student.Name]; exists {
		return fmt.Errorf("%s: %w", student.Name, errors.ErrAlreadyExists)
	}

	students[student.Name] = student

	return nil
}

func (s *Storage) CreateSubject(_ context.Context, sub server.Subject) (err error) {
	sub.ID = uuid.New()
	subjects[sub.ID] = sub
	return nil
}

func (s *Storage) CreateTheme(_ context.Context, theme server.Theme) (err error) {
	theme.ID = uuid.New()
	themes[theme.ID] = theme
	return nil
}

func (s *Storage) GetStudent(_ context.Context, name string) (student server.Student, err error) {
	var exists bool
	student, exists = students[name]
	if !exists {
		return student, fmt.Errorf("%s: %w", name, errors.ErrNotFound)
	}

	return student, nil
}

func (s *Storage) ListThemes(_ context.Context, student string) (res []server.Theme, err error) {
	for _, theme := range themes {
		if theme.UserName == student {
			res = append(res, theme)
		}
	}

	return res, nil
}
