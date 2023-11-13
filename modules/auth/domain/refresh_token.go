package domain

import pbUser "github.com/labbs/castle/gen/user"

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByUsername(username string) (pbUser.User, error)
	CreateAccessToken(user *pbUser.User) (accessToken string, err error)
	CreateRefreshToken(user *pbUser.User) (refreshToken string, err error)
	GetIdFromToken(token string) (string, error)
}
