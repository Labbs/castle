package repository

import (
	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/internal/bus"
	"github.com/labbs/castle/modules/auth/internal/tokenutil"
)

type AuthRepository struct {
	BusMessages chan initBootstrap.Message
}

func NewAuthRepository(busMessages chan initBootstrap.Message) AuthRepository {
	return AuthRepository{BusMessages: busMessages}
}

func (d *AuthRepository) GetUserByUsername(username string) (bus.GetUserByUsernameResponse, error) {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "user:get_by_username", Data: username, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return bus.GetUserByUsernameResponse{}, response.Error
	}
	return response.Data.(bus.GetUserByUsernameResponse), nil
}

func (d *AuthRepository) CreateAccessToken(username string) (string, error) {
	return tokenutil.CreateAccessToken(username)
}

func (d *AuthRepository) CreateRefreshToken(username string) (string, error) {
	return tokenutil.CreateRefreshToken(username)
}

func (d *AuthRepository) GetUsernameFromToken(token string) (string, error) {
	return tokenutil.GetUsernameFromToken(token)
}
