package server

import (
	"github.com/Henrod/study-track/studytrack"
	"github.com/google/uuid"
)

type Theme struct {
	ID       uuid.UUID
	UserName string // fk
	Name     string
}

func (*Theme) fromProto(p *studytrack.Theme, studentName string) *Theme {
	return &Theme{
		ID:       uuid.Nil,
		UserName: studentName,
		Name:     p.Name,
	}
}

func (t *Theme) toProto() *studytrack.Theme {
	return &studytrack.Theme{
		Name: t.Name,
	}
}

type Themes []Theme

func (t Themes) toProto() []*studytrack.Theme {
	p := make([]*studytrack.Theme, 0, len(t))
	for _, theme := range t {
		p = append(p, theme.toProto())
	}
	return p
}
