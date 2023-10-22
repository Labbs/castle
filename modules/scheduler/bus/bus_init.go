package bus

import (
	"github.com/labbs/castle/modules/scheduler/bootstrap"
	"github.com/labbs/castle/modules/scheduler/repository"
)

type Controller struct {
	Repository repository.CronRepository
}

func Setup(app bootstrap.Application) {
	ur := repository.NewCronRepository(app.CronScheduler)
	uc := &Controller{
		Repository: ur,
	}

	app.Bus.AddHandler("scheduler:cron:add", uc.AddSchedulerCronJob)
}
