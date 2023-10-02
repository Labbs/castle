package bus

import "github.com/labbs/castle/internal/bus"

func (uc *UserController) GetByUsername(data interface{}) interface{} {
	username := data.(string)
	user, err := uc.Repository.GetUserByUsername(username)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return bus.GetUserByUsernameResponse{Username: user.Username, Password: user.Password}
}
