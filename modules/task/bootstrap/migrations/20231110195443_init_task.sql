-- +goose Up
-- +goose StatementBegin
CREATE TABLE task (
  id text PRIMARY KEY,
  name text NOT NULL,
  description text,

  project_id text NOT NULL,
  environment_id text NOT NULL,
  variables jsonb NOT NULL DEFAULT '{}',
  repository_id text NOT NULL,
  repository_path text NOT NULL,
  enabled boolean NOT NULL,
  task_type int NOT NULL,
  exec_type int NOT NULL,
  ansible_task jsonb NOT NULL DEFAULT '{}',
  terraform_task jsonb NOT NULL DEFAULT '{}',
  script_task jsonb NOT NULL DEFAULT '{}',

  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task;
-- +goose StatementEnd
