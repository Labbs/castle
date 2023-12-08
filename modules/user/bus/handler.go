package bus

import (
	"github.com/goccy/go-json"
)

func (uc *UserController) GetByEmail(data interface{}) (interface{}, error) {
	email := data.(string)
	uc.Logger.Debug().Str("event", "bus.user.get_by_email").Str("email", email).Msg("requesting user from db")
	user, err := uc.Repository.GetUserByEmail(email)
	if err != nil {
		uc.Logger.Debug().Err(err).Str("event", "bus.user.get_by_email").Str("email", email).Msg("failed to get user from db")
		return map[string]string{}, err
	}
	uc.Logger.Debug().Str("event", "bus.user.get_by_email").Interface("user", user).Msg("user found")

	j, err := json.Marshal(user)
	if err != nil {
		uc.Logger.Debug().Err(err).Str("event", "bus.user.get_by_email").Interface("user", user).Msg("failed to marshal user")
		return map[string]string{}, err
	}
	uc.Logger.Debug().Str("event", "bus.user.get_by_email").Interface("user", user).Msg("user marshalled")

	return j, nil
}
