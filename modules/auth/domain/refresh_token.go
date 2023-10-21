package domain

import userDomain "github.com/labbs/castle/modules/user/domain"

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByUsername(username string) (userDomain.User, error)
	CreateAccessToken(user *userDomain.User) (accessToken string, err error)
	CreateRefreshToken(user *userDomain.User) (refreshToken string, err error)
	GetIdFromToken(token string) (string, error)
}
