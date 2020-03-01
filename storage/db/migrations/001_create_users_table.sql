-- +migrate Up
CREATE TABLE users (
  id   UUID PRIMARY KEY,
  name TEXT NOT NULL
);

-- +migrate Down
DROP TABLE users;
