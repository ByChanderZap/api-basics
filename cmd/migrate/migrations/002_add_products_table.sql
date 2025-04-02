-- +goose Up
CREATE TABLE IF NOT EXISTS products (
  id              UUID PRIMARY KEY,
  name            TEXT NOT NULL,
  description     TEXT NOT NULL,
  image           TEXT NULL,
  price           DECIMAL(10, 2) NOT NULL,
  quantity        INT NOT NULL,

  created_at      TIMESTAMP NOT NULL,
  updated_at      TIMESTAMP NOT NULL,
  deleted_at      TIMESTAMP NULL
);
--

-- +goose Down
DROP TABLE IF EXISTS products;
--
