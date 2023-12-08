package bus

import (
	"github.com/labbs/castle/modules/task/bootstrap"
	"github.com/labbs/castle/modules/task/domain"
	"github.com/labbs/castle/modules/task/repository"
	"github.com/labbs/castle/modules/task/service"
	"github.com/rs/zerolog"
)

type Controller struct {
	Service domain.TaskService
	Logger  zerolog.Logger
}

func Setup(app bootstrap.Application) {
	ur := repository.NewTaskRepository(app.Db)
	uc := &Controller{
		Service: service.NewTaskService(ur),
		Logger:  app.Logger,
	}

	app.Bus.AddHandler("task:get_by_id", uc.GetTaskById)
	app.Bus.AddHandler("task:get_all", uc.GetAllTasks)
	app.Bus.AddHandler("task:get_all_by_project_id", uc.GetAllTasksByProjectId)
	app.Bus.AddHandler("task:create", uc.CreateTask)
	app.Bus.AddHandler("task:count_by_repository_id", uc.CountTasksByRepositoryId)
}
