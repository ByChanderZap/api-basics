-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id            UUID PRIMARY KEY,
  firstName     TEXT NOT NULL,
  lastName      TEXT NOT NULL,
  email         TEXT UNIQUE NOT NULL,
  password      TEXT NOT NULL,
  
  createdAt     TIMESTAMP NOT NULL,
  updatedAt     TIMESTAMP NOT NULL,
  deletedAt     TIMESTAMP
);
--

-- +goose Down
DROP TABLE IF EXISTS users;
--
