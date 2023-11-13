package bus

import "github.com/goccy/go-json"

func (uc *UserController) GetByEmail(data interface{}) interface{} {
	email := data.(string)
	user, err := uc.Repository.GetUserByUsername(email)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	j, err := json.Marshal(user)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return j
}
