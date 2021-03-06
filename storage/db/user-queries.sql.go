// Code generated by sqlc. DO NOT EDIT.
// source: user-queries.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createSubject = `-- name: CreateSubject :one
INSERT INTO subjects (
    user_id,
    title,
    description,
    deadline,
    created_at,
    updated_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    now(),
    now()
) RETURNING id, user_id, title, description, deadline, created_at, updated_at
`

type CreateSubjectParams struct {
	UserID      uuid.UUID    `json:"user_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Deadline    sql.NullTime `json:"deadline"`
}

func (q *Queries) CreateSubject(ctx context.Context, arg CreateSubjectParams) (Subject, error) {
	row := q.db.QueryRowContext(ctx, createSubject,
		arg.UserID,
		arg.Title,
		arg.Description,
		arg.Deadline,
	)
	var i Subject
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Deadline,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, name) VALUES ($1, $2) RETURNING id, name
`

type CreateUserParams struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.Name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
