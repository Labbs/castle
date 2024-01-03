package bus

import (
	"github.com/labbs/castle/modules/scheduler/bootstrap"
	"github.com/labbs/castle/modules/scheduler/domain"
	"github.com/labbs/castle/modules/scheduler/repository"
	"github.com/labbs/castle/modules/scheduler/service"
	"github.com/rs/zerolog"
)

type SchedulerController struct {
	Service domain.SchedulerService
	Logger  zerolog.Logger
}

func Setup(app bootstrap.Application) {
	ur := repository.NewSchedulerRepository(app.CronScheduler, app.Logger)
	uc := &SchedulerController{
		Service: service.NewSchedulerService(ur),
		Logger:  app.Logger,
	}

	app.Bus.AddHandler("scheduler:cron:add", uc.AddSchedulerCronJob)
	app.Bus.AddHandler("scheduler:cron:remove", uc.RemoveSchedulerCronJob)
	app.Bus.AddHandler("scheduler:cron:get_all", uc.GetAllSchedulerCronJobs)
	app.Bus.AddHandler("scheduler:cron:run", uc.RunSchedulerCronJob)
}
