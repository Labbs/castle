-- +goose Up
-- +goose StatementBegin
CREATE TABLE repository (
  id text PRIMARY KEY,
  name text NOT NULL,
  description text,
  type int NOT NULL,
  url text NOT NULL,

  ssh_key text,
  ssh_key_passphrase text,

  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE repository;
-- +goose StatementEnd
