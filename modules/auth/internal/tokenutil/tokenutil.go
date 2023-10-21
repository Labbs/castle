package tokenutil

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labbs/castle/config"
	"github.com/labbs/castle/modules/auth/domain"
)

func CreateAccessToken(username string) (accessToken string, err error) {
	exp := time.Now().Add(time.Second * time.Duration(config.AppConfig.Session.Expire)).Unix()
	claims := &domain.JwtCustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.AppConfig.Session.Issuer,
			ExpiresAt: &jwt.NumericDate{Time: time.Unix(exp, 0)},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.AppConfig.Session.SecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func CreateRefreshToken(username string) (refreshToken string, err error) {
	exp := time.Now().Add(time.Second * time.Duration(config.AppConfig.Session.Expire)).Unix()
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.AppConfig.Session.Issuer,
			ExpiresAt: &jwt.NumericDate{Time: time.Unix(exp, 0)},
		},
	}
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	t, err := tokenRefresh.SignedString([]byte(config.AppConfig.Session.SecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetUsernameFromToken(token string) (username string, err error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.Session.SecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		username = claims["username"].(string)
		return username, nil
	}

	return "", fmt.Errorf("invalid token")
}

func IsAuthorized(token string) (bool, error) {
	_, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.Session.SecretKey), nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
