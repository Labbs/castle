package repository

import (
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"

	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/modules/auth/domain"
	"github.com/labbs/castle/modules/auth/internal/tokenutil"
)

type AuthRepository struct {
	BusMessages chan initBootstrap.Message
	Logger      zerolog.Logger
}

func NewAuthRepository(busMessages chan initBootstrap.Message, logger zerolog.Logger) AuthRepository {
	return AuthRepository{BusMessages: busMessages, Logger: logger}
}

func (d *AuthRepository) GetUserByEmail(email string) (domain.BusGetUserByEmailResponse, error) {
	d.Logger.Debug().Str("event", "repository.auth.get_user_by_email").Str("email", email).Msg("requesting user from bus")

	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "user:get_by_email", Data: email, Response: responseChan}
	response := <-responseChan
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

func (d *AuthRepository) CreateAccessToken(email string) (string, error) {
	return tokenutil.CreateAccessToken(email)
}

func (d *AuthRepository) CreateRefreshToken(email string) (string, error) {
	return tokenutil.CreateRefreshToken(email)
}

func (d *AuthRepository) GetEmailFromToken(token string) (string, error) {
	return tokenutil.GetEmailFromToken(token)
}
