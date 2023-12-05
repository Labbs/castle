package repository

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/rs/zerolog"

	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/modules/frontend/domain"
)

type UserRepository struct {
	BusMessages chan initBootstrap.Message
	Logger      zerolog.Logger
}

func NewUserRepository(busMessages chan initBootstrap.Message, logger zerolog.Logger) UserRepository {
	return UserRepository{BusMessages: busMessages, Logger: logger}
}

func (d *UserRepository) GetUserByEmail(email string) (domain.BusGetUserByEmailResponse, error) {
	d.Logger.Debug().Str("event", "repository.auth.get_user_by_email").Str("email", email).Msg("requesting user from bus")

	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "user:get_by_email", Data: email, Response: responseChan}
	response := <-responseChan
	d.Logger.Debug().Str("event", "repository.auth.get_user_by_email").Interface("response", response).Msg("response received")
	fmt.Println(response)
	if response.Error != nil {
		return domain.BusGetUserByEmailResponse{}, response.Error
	}
	d.Logger.Debug().Str("event", "repository.auth.get_user_by_email").Interface("user", response.Data).Msg("user found")

	var user domain.BusGetUserByEmailResponse
	err := json.Unmarshal(response.Data.([]byte), &user)
	if err != nil {
		return domain.BusGetUserByEmailResponse{}, err
	}
	d.Logger.Debug().Str("event", "repository.auth.get_user_by_email").Str("email", email).Msg("user unmarshalled")

	return user, nil
}
