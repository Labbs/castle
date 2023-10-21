package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/modules/user/domain"
	"github.com/labbs/castle/modules/user/repository"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Repository repository.UserRepository
	Logger     zerolog.Logger
}

func (pc *UserController) Get(c *fiber.Ctx) error {
	user, err := pc.Repository.GetUserByUsername(c.Context().UserValue("username").(string))
	if err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.get").Msg("failed to find user with username")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get profile", "status": "error"})
	}

	user.Password = ""
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func (pc *UserController) EditUsername(c *fiber.Ctx) error {
	user, err := pc.Repository.GetUserByUsername(c.Context().UserValue("username").(string))
	if err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_username").Msg("failed to find user with username")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit username", "status": "error"})
	}

	request := new(domain.UsernameChangeRequest)
	if err := c.BodyParser(request); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_username").Msg("failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request", "status": "error"})
	}

	if user.Username != request.CurrentUsername {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Current username is incorrect", "status": "error"})
	}

	user.Username = request.NewUsername
	if err := pc.Repository.UpdateUser(user); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_username").Msg("failed to update user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit username", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Username updated successfully",
		"status":  "success",
	})
}

func (pc *UserController) EditAvatar(c *fiber.Ctx) error {
	user, err := pc.Repository.GetUserByUsername(c.Context().UserValue("username").(string))
	if err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_avatar").Msg("failed to find user with username")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit avatar", "status": "error"})
	}

	request := new(domain.AvatarChangeRequest)
	if err := c.BodyParser(request); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_avatar").Msg("failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request", "status": "error"})
	}

	user.Avatar = request.Avatar
	if err := pc.Repository.UpdateUser(user); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_avatar").Msg("failed to update user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit avatar", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Avatar updated successfully",
		"status":  "success",
	})
}

func (pc *UserController) EditDarkMode(c *fiber.Ctx) error {
	user, err := pc.Repository.GetUserByUsername(c.Context().UserValue("username").(string))
	if err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_dark_mode").Msg("failed to find user with username")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit dark mode", "status": "error"})
	}

	request := new(domain.DarkModeChangeRequest)
	if err := c.BodyParser(request); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_dark_mode").Msg("failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request", "status": "error"})
	}

	user.DarkMode = request.DarkMode
	if err := pc.Repository.UpdateUser(user); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_dark_mode").Msg("failed to update user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit dark mode", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Dark mode updated successfully",
		"status":  "success",
	})
}

func (pc *UserController) EditPassword(c *fiber.Ctx) error {
	user, err := pc.Repository.GetUserByUsername(c.Context().UserValue("username").(string))
	if err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_password").Msg("failed to find user with username")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit password", "status": "error"})
	}

	request := new(domain.PasswordChangeRequest)
	if err := c.BodyParser(request); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_password").Msg("failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request", "status": "error"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.CurrentPassword)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Current password is incorrect", "status": "error"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 30)
	if err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_password").Msg("failed to hash password")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit password", "status": "error"})
	}

	user.Password = string(hashedPassword)

	if err := pc.Repository.UpdateUser(user); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.profile.edit_password").Msg("failed to update user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit password", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password updated successfully",
		"status":  "success",
	})
}
