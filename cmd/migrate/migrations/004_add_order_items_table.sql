-- +goose Up
CREATE TABLE IF NOT EXISTS order_items (
  id                UUID PRIMARY KEY,
  order_id          UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  product_id        UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
  quantity          INT NOT NULL,
  price             DECIMAL(10, 2) NOT NULL
);
--

-- +goose Down
DROP TABLE IF EXISTS order_items;
--
