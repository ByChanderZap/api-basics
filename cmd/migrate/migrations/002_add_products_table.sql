-- +goose Up
CREATE TABLE IF NOT EXISTS products (
  id            UUID PRIMARY KEY,
  name          TEXT NOT NULL,
  description   TEXT NOT NULL,
  image         TEXT NOT NULL,
  price         DECIMAL(10, 2) NOT NULL,
  quantity      INT NOT NULL,

  createdAt     TIMESTAMP NOT NULL,
  updatedAt     TIMESTAMP NOT NULL,
  deletedAt     TIMESTAMP
);
--

-- +goose Down
DROP TABLE IF EXISTS products;
--
