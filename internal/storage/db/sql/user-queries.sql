CREATE TABLE users (
  id   UUID PRIMARY KEY,
  name TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- name: CreateUser :one
INSERT INTO users (id, name) VALUES ($1, $2) RETURNING *;

CREATE TABLE users (
    id   UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    deadline TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- name: CreateSubject :one
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
) RETURNING *;
