package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/labbs/castle/config"
	"github.com/labbs/castle/internal"
	"github.com/labbs/castle/modules/frontend/repository"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	Repository repository.UserRepository
	Logger     zerolog.Logger
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	d := make(fiber.Map)
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	if config.AppConfig.LocalDev {
		d["LocalDev"] = true
	}

	if c.Method() == "POST" {
		user, err := ac.Repository.GetUserByUsername(c.FormValue("username"))
		if err != nil {
			d["Error"] = "Invalid username or password"
			return c.Render("templates/login", d)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.FormValue("password")))
		if err != nil {
			d["Error"] = "Invalid username or password"
			return c.Render("templates/login", d)
		}

		store, _ := c.Locals("sessions").(*session.Store).Get(c)
		store.Set("username", user.Username)
		user.Password = ""
		m, _ := internal.StructToMap(user)
		store.Set("profile", m)

		err = store.Save()
		if err != nil {
			ac.Logger.Error().Err(err).Str("event", "frontend.controller.auth.login").Msg("failed to save session")
			d["Error"] = "Internal server error"
			return c.Render("templates/login", d)
		}

		return c.Redirect("/app")
	}

	return c.Render("templates/login", d)
}

func (ac *AuthController) Logout(c *fiber.Ctx) error {
	store, _ := c.Locals("sessions").(*session.Store).Get(c)
	store.Delete("username")
	store.Delete("profile")
	err := store.Destroy()
	if err != nil {
		ac.Logger.Error().Err(err).Str("event", "frontend.controller.auth.logout").Msg("failed to destroy session")
	}

	return c.Redirect("/app/auth/login")
}
