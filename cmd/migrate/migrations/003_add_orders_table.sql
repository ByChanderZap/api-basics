-- +goose Up
CREATE TYPE order_status AS ENUM ('pending', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
  id            UUID PRIMARY KEY,
  userId        UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  total         DECIMAL(10, 2) NOT NULL,
  status        order_status NOT NULL DEFAULT 'pending',
  address       TEXT NOT NULL,
  createdAt     TIMESTAMP NOT NULL,
  updatedAt     TIMESTAMP NOT NULL,
  deletedAt     TIMESTAMP
);
--

-- +goose Down
DROP TABLE IF EXISTS orders;
DROP TYPE IF EXISTS order_status;
--
