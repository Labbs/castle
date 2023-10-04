package bus

func (uc *RepositoryController) GetRepositoryById(data interface{}) interface{} {
	id := data.(string)
	repo, err := uc.Repository.GetRepositoryById(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return repo
}
