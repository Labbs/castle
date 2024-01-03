package bus

import (
	"github.com/goccy/go-json"

	"github.com/labbs/castle/modules/scheduler/domain"
)

// Bus handler for scheduler:cron:add
func (uc *SchedulerController) AddSchedulerCronJob(data interface{}) (interface{}, error) {
	var task domain.Task
	err := json.Unmarshal(data.([]byte), &task)
	if err != nil {
		return map[string]string{}, err
	}

	err = uc.Service.AddSchedulerCronJob(task)
	if err != nil {
		return map[string]string{}, err
	}
	return map[string]string{"success": "scheduler cron job added"}, nil
}

// Bus handler for scheduler:cron:remove
func (uc *SchedulerController) RemoveSchedulerCronJob(data interface{}) (interface{}, error) {
	jobId := data.(string)
	err := uc.Service.RemoveSchedulerCronJob(jobId)
	if err != nil {
		return map[string]string{}, err
	}
	return map[string]string{"success": "scheduler cron job removed"}, nil
}

// Bus handler for scheduler:cron:get_all
func (uc *SchedulerController) GetAllSchedulerCronJobs(data interface{}) (interface{}, error) {
	jobs, err := uc.Service.GetAllSchedulerCronJobs()
	if err != nil {
		return map[string]string{}, err
	}

	j, err := json.Marshal(jobs)
	if err != nil {
		return map[string]string{}, err
	}
	return j, nil
}

// Bus handler for scheduler:cron:run
func (uc *SchedulerController) RunSchedulerCronJob(data interface{}) (interface{}, error) {
	return nil, nil
}
