package repository

import (
	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/modules/auth/domain"
	"github.com/labbs/castle/modules/auth/internal/tokenutil"
)

type AuthRepository struct {
	BusMessages chan initBootstrap.Message
}

func NewAuthRepository(busMessages chan initBootstrap.Message) AuthRepository {
	return AuthRepository{BusMessages: busMessages}
}

func (d *AuthRepository) GetUserByUsername(username string) (domain.BusGetUserByUsernameResponse, error) {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "user:get_by_username", Data: username, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return domain.BusGetUserByUsernameResponse{}, response.Error
	}
	return response.Data.(domain.BusGetUserByUsernameResponse), nil
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
