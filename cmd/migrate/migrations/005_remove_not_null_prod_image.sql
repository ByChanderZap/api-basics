-- +goose Up
-- +goose StatementBegin
ALTER TABLE products ALTER COLUMN image DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products ALTER COLUMN image SET NOT NULL;
-- +goose StatementEnd
