package bus

func (uc *Controller) GetTaskById(data interface{}) interface{} {
	task := data.(string)
	r, err := uc.Repository.GetTaskById(task)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return r
}

func (uc *Controller) GetAllTasks(data interface{}) interface{} {
	r, err := uc.Repository.GetAllTasks()
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return r
}
