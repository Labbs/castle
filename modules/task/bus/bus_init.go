package bus

import (
	"github.com/labbs/castle/modules/task/bootstrap"
	"github.com/labbs/castle/modules/task/repository"
)

type Controller struct {
	Repository repository.TaskRepository
}

func Setup(app bootstrap.Application) {
	ur := repository.NewTaskRepository(app.Db)
	uc := &Controller{
		Repository: ur,
	}

	app.Bus.AddHandler("task:get_by_id", uc.GetTaskById)
	app.Bus.AddHandler("task:get_all", uc.GetAllTasks)
	app.Bus.AddHandler("task:get_all_by_project_id", uc.GetAllTasksByProjectId)
	app.Bus.AddHandler("task:create", uc.CreateTask)
	app.Bus.AddHandler("task:count_by_repository_id", uc.CountTasksByRepositoryId)
}
