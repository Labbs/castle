# Castle

## Description
Castle is a scheduler for execute Terraform/Ansible/Shell/... scripts from a repository.

## Development
- Language: Golang 1.21
- Framework: Fiber, Gorm, ZeroLog, Urfave Cli

### The organization of the code
- bootstrap: functions for initialize the database, http server, logger, ...
- cmd: the main function
- config: the configuration of the application
- internal: common functions for the whole application
- modules: the modules of the application
- modules/auth: the module for authentication (Very simple authentication with a token and use the user module for get the user information)
- modules/frontend: the module for the frontend (HTML, CSS, JS, ... with golang template and fiber. The frontend is embedded in the binary)
- modules/notification: the module for the notification (Slack, Discord, ...) - Not implemented
- modules/user: the module for the user (user, group, role, permission, ...) - Group, role and permission are not implemented
- modules/project: the module for the project (project, environment, ...) - Environment is not implemented
- modules/task: the module for the task (task, task type, task status, task log, ...) - Task status and task log are not implemented
- .dev-resources: the resources for the development (http requests for populate the database, ...)

### Generate the PB models

```bash
### Install the protoc compiler
brew install protobuf

### Install the protoc go plugin
make install-pb-deps

### Generate the PB models
make generate-pb
```

## License
Apache License 2.0