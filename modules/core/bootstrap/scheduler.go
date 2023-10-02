package bootstrap

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/labbs/castle/modules/core/internal/scheduler"
	"github.com/labbs/castle/modules/core/repository"
)

func InitScheduler(app Application) scheduler.SchedulerController {
	cronScheduler := gocron.NewScheduler(time.UTC)
	cronScheduler.TagsUnique()
	cronScheduler.StartAsync()

	return scheduler.SchedulerController{
		Repository: repository.NewCoreRepository(app.Db),
		Logger:     app.Logger,
		Scheduler:  cronScheduler,
	}
}
