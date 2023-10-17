package bootstrap

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/labbs/castle/config"
	"github.com/labbs/castle/modules/frontend/repository"
)

func InitSession(app *Application) {
	session := session.New(session.Config{
		Expiration: time.Second * time.Duration(config.AppConfig.Session.Expire),
	})

	session.RegisterType(map[string]interface{}{})
	session.RegisterType([]interface{}{})

	app.Fiber.Use(func(c *fiber.Ctx) error {
		c.Locals("sessions", session)
		return c.Next()
	})

	app.Fiber.Use(checkSession())

	app.Fiber.Use(func(c *fiber.Ctx) error {
		if c.Path() != "/app/auth/login" {
			c.Locals("latestUserData", getLatestUserData(c, app))
		}
		return c.Next()
	})
}

func checkSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		store, err := c.Locals("sessions").(*session.Store).Get(c)
		if err != nil {
			return c.Next()
		}

		if c.Path() != "/app/auth/login" {
			if store.Get("username") == nil {
				return c.Redirect("/app/auth/login", fiber.StatusTemporaryRedirect)
			}
		}

		return c.Next()
	}
}

func getLatestUserData(c *fiber.Ctx, app *Application) fiber.Map {
	d := make(fiber.Map)

	if config.AppConfig.LocalDev {
		d["LocalDev"] = true
	}

	store, _ := c.Locals("sessions").(*session.Store).Get(c)

	r := repository.NewUserRepository(app.BusMessages)
	user, err := r.GetUserByUsername(store.Get("username").(string))
	if err != nil {
		app.Logger.Error().Err(err).Str("event", "frontend.session.get_latest_user_data").Msg("failed to get user data")
		return d
	}

	user.Password = ""
	d["Profile"] = user

	return d
}
