package server

import (
	"time"

	"github.com/Henrod/study-track/studytrack"
)

type Student struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*Student) fromProto(p *studytrack.Student) *Student {
	return &Student{
		Name: p.Name,
	}
}

func (s *Student) toProto() *studytrack.Student {
	return &studytrack.Student{
		Name: s.Name,
	}
}
