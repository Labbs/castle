-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
  id text PRIMARY KEY,
  email text NOT NULL,
  password text NOT NULL,
  avatar text,
  dark_mode int NOT NULL DEFAULT 2,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd
