-- +goose Up
-- +goose StatementBegin
CREATE TABLE environment (
  id text PRIMARY KEY,
  name text NOT NULL,
  description text,

  project_id text NOT NULL REFERENCES project(id),
  variables jsonb NOT NULL DEFAULT '{}',

  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE environment;
-- +goose StatementEnd
