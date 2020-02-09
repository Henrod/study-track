package storage

import (
	"time"

	"github.com/Henrod/study-track/internal/bll"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Subject struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Title       string
	Description string
	Deadline    pq.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func FromBLLSubject(s bll.Subject) Subject {
	return Subject{
		ID:          s.ID,
		UserID:      s.UserID,
		Title:       s.Title,
		Description: s.Description,
		Deadline: pq.NullTime{
			Time:  s.Deadline,
			Valid: s.Deadline.IsZero(),
		},
	}
}
