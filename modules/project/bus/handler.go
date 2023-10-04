package bus

func (uc *Controller) GetProjectById(data interface{}) interface{} {
	id := data.(string)
	r, err := uc.Repository.GetProjectById(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return r
}

func (uc *Controller) GetAllProjects(data interface{}) interface{} {
	r, err := uc.Repository.GetAllProjects()
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return r
}
