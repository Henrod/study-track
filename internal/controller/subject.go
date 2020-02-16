package controller

import (
	"context"

	"github.com/Henrod/study-track/internal/bll"

	"github.com/Henrod/study-track/internal/errors"

	"github.com/Henrod/study-track/pkg/studytrack"
	"github.com/sirupsen/logrus"
)

type Subject struct {
	storage bll.Storage
	logger  logrus.FieldLogger
}

var _ studytrack.SubjectServiceServer = &Subject{}

func NewSubject(storage bll.Storage, logger logrus.FieldLogger) *Subject {
	return &Subject{storage: storage, logger: logger}
}

func (u *Subject) Create(
	_ context.Context,
	_ *studytrack.CreateSubjectRequest,
) (res *studytrack.Subject, err error) {
	return res, errors.ErrMissingImplementation
}
