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
}
