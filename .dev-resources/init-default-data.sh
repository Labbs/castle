#!/bin/bash

set -x
# Init default values for the local development environment

# default values for the local development environment
API_USER="default"
API_PASSWORD="localdev"
API_PORT="8080"
API_HOST="localhost"
API_PROTOCOL="http"

DEFAULT_REPOSITORY_NAME="castle-repo-example"
DEFAULT_REPOSITORY_URL="git@github.com:Labbs/castle-repo-example.git"

DEFAULT_PROJECT_NAME="example"

# authentication to the api with curl
TOKEN=`curl -s -XPOST -H "Content-Type: application/json" -d '{"username": "'${API_USER}'", "password": "'${API_PASSWORD}'"}' $API_PROTOCOL://$API_HOST:$API_PORT/api/auth/login | jq -r '.access_token'`

# create a repository
REPOSITORY_ID=`curl -s -XPOST -H "Content-Type: application/json" -H "Authorization: Bearer ${TOKEN}" \
  -d '{"name": "'${DEFAULT_REPOSITORY_NAME}'", "url": "'${DEFAULT_REPOSITORY_URL}'"}' \
  $API_PROTOCOL://$API_HOST:$API_PORT/api/repository/create | jq -r '.repository.id'`

echo "Repository created with id: ${REPOSITORY_ID}"

# create a project
PROJECT_ID=`curl -s -XPOST -H "Content-Type: application/json" -H "Authorization: Bearer ${TOKEN}" \
  -d '{"name": "'${DEFAULT_PROJECT_NAME}'"}' \
  $API_PROTOCOL://$API_HOST:$API_PORT/api/project/create | jq -r '.project.id'`

echo "Project created with id: ${PROJECT_ID}"

# create a task
TASK_ID=`curl -s -XPOST -H "Content-Type: application/json" -H "Authorization: Bearer ${TOKEN}" \
  -d '{"name": "example", "repository_id": "'${REPOSITORY_ID}'", "project_id": "'${PROJECT_ID}', "enabled": "false", "repository_path": "terraform/simple-example", "task_type": "terraform"}' \
  $API_PROTOCOL://$API_HOST:$API_PORT/api/task/create | jq -r '.task.id'`

echo "Task created with id: ${TASK_ID}"