package server

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/Henrod/study-track/studytrack"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Subject struct {
	ID          uuid.UUID
	ThemeID     uuid.UUID // fk
	Title       string
	Description string
	Deadline    pq.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (*Subject) fromProto(p *studytrack.Subject) *Subject {
	return &Subject{
		ID:          uuid.Nil,
		ThemeID:     uuid.Nil,
		Title:       p.Title,
		Description: p.Description,
		Deadline:    pq.NullTime{Time: time.Unix(p.Deadline.Seconds, 0), Valid: true},
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func (s *Subject) toProto() *studytrack.Subject {
	return &studytrack.Subject{
		Title:       s.Title,
		Description: s.Description,
		Deadline:    &timestamp.Timestamp{Seconds: s.Deadline.Time.Unix()},
	}
}
