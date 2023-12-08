// File: login_controller.go
// Code Responsible for handling login requests
package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/modules/auth/domain"
	"github.com/labbs/castle/modules/auth/repository"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

// AuthController struct
type AuthController struct {
	Repository repository.AuthRepository
	Logger     zerolog.Logger
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	ac.Logger.Debug().Str("event", "api.controller.auth.login").Msg("login request received")

	request := new(domain.LoginRequest)
	if err := c.BodyParser(request); err != nil {
		ac.Logger.Error().Err(err).Str("event", "api.controller.auth.login").Msg("failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	ac.Logger.Debug().Str("event", "api.controller.auth.login").Interface("request", request).Msg("request parsed successfully")

	user, err := ac.Repository.GetUserByEmail(request.Email)
	if err != nil {
		ac.Logger.Error().Err(err).Str("event", "api.controller.auth.login").Msg("failed to find user")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email or password"})
	}
	ac.Logger.Debug().Str("event", "api.controller.auth.login").Msg("user found")

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		ac.Logger.Error().Err(err).Str("event", "api.controller.auth.login").Msg("failed to compare password")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email or password"})
	}
	ac.Logger.Debug().Str("event", "api.controller.auth.login").Msg("password matched")

	accessToken, err := ac.Repository.CreateAccessToken(user.Email)
	if err != nil {
		ac.Logger.Error().Err(err).Str("event", "api.controller.auth.login").Msg("failed to create access token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	ac.Logger.Debug().Str("event", "api.controller.auth.login").Msg("access token created")

	refreshToken, err := ac.Repository.CreateRefreshToken(user.Email)
	if err != nil {
		ac.Logger.Error().Err(err).Str("event", "api.controller.auth.login").Msg("failed to create refresh token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	ac.Logger.Debug().Str("event", "api.controller.auth.login").Msg("refresh token created")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (ac *AuthController) RefreshToken(c *fiber.Ctx) error {
	request := new(domain.RefreshTokenRequest)
	if err := c.BodyParser(request); err != nil {
		ac.Logger.Error().Err(err).Str("event", "api.controller.auth.refresh_token").Msg("failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	userId, err := ac.Repository.GetEmailFromToken(request.RefreshToken)
	if err != nil {
		ac.Logger.Error().Err(err).Str("event", "api.controller.auth.refresh_token").Msg("failed to parse refresh token")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid refresh token"})
	}

	user, err := ac.Repository.GetUserByEmail(userId)
	if err != nil {
		ac.Logger.Error().Err(err).Str("event", "api.controller.auth.refresh_token").Msg("failed to find user")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid refresh token"})
	}

	accessToken, err := ac.Repository.CreateAccessToken(user.Email)
	if err != nil {
		ac.Logger.Error().Err(err).Str("event", "api.controller.auth.refresh_token").Msg("failed to create access token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	refreshToken, err := ac.Repository.CreateRefreshToken(user.Email)
	if err != nil {
		ac.Logger.Error().Err(err).Str("event", "api.controller.auth.refresh_token").Msg("failed to create refresh token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
