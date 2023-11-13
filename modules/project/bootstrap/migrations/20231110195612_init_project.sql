-- +goose Up
-- +goose StatementBegin
CREATE TABLE project (
  id text PRIMARY KEY,
  name text NOT NULL,
  description text,

  variables jsonb NOT NULL DEFAULT '{}',

  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE project;
-- +goose StatementEnd
