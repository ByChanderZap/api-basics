-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id              UUID PRIMARY KEY,
  first_name      TEXT NOT NULL,
  last_name       TEXT NOT NULL,
  email           TEXT UNIQUE NOT NULL,
  password        TEXT NOT NULL,
  
  created_at      TIMESTAMP NOT NULL,
  updated_at      TIMESTAMP NOT NULL,
  deleted_at      TIMESTAMP
);
--

-- +goose Down
DROP TABLE IF EXISTS users;
--
