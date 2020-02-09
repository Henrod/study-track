package bll

import (
	"context"
	"fmt"
	"time"

	"github.com/Henrod/study-track/internal/errors"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

type Subject struct {
	ID          uuid.UUID
	SubjectID   uuid.UUID
	Title       string
	Description string
	Deadline    time.Time
}

func CreateSubject(
	ctx context.Context,
	logger logrus.FieldLogger,
	storage Storage,
	subject Subject,
) (err error) {
	err = storage.CreateSubject(ctx, subject)
	if err != nil {
		logger.WithError(err).Errorf("failed to create subject")
		return fmt.Errorf("failed to create subject: %w", errors.ErrInternal)
	}

	return nil
}
