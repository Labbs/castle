package repository

import (
	"github.com/goccy/go-json"

	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/modules/frontend/domain"
)

type UserRepository struct {
	BusMessages chan initBootstrap.Message
}

func NewUserRepository(busMessages chan initBootstrap.Message) UserRepository {
	return UserRepository{BusMessages: busMessages}
}

func (d *UserRepository) GetUserByUsername(username string) (domain.BusGetUserByUsernameResponse, error) {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "user:get_by_username", Data: username, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return domain.BusGetUserByUsernameResponse{}, response.Error
	}

	var user domain.BusGetUserByUsernameResponse
	err := json.Unmarshal(response.Data.([]byte), &user)
	if err != nil {
		return domain.BusGetUserByUsernameResponse{}, err
	}

	return user, nil
}
