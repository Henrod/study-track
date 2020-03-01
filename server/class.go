package server

import (
	"context"
	"strings"

	"github.com/Henrod/study-track/studytrack"
	"github.com/sirupsen/logrus"
)

// Class contain student related functions
type Class struct {
	storage Storage
	logger  logrus.FieldLogger
}

func NewClass(storage Storage, logger logrus.FieldLogger) *Class {
	return &Class{storage: storage, logger: logger}
}

var _ studytrack.ClassServer = &Class{}

func (c *Class) CreateStudent(
	ctx context.Context,
	req *studytrack.CreateStudentRequest,
) (res *studytrack.Student, err error) {
	student := (*Student)(nil).fromProto(req.Student)

	err = c.storage.CreateStudent(ctx, *student)
	if err != nil {
		return nil, parseErr(err, "failed to create student")
	}

	return student.toProto(), nil
}

func (c *Class) GetStudent(
	ctx context.Context,
	req *studytrack.GetStudentRequest,
) (res *studytrack.Student, err error) {
	name := strings.Split(req.Name, "/")[1]

	student, err := c.storage.GetStudent(ctx, name)
	if err != nil {
		return nil, parseErr(err, "failed to get student")
	}

	return student.toProto(), nil
}
