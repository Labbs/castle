// Path: bootstrap/fiber.go
// This is a generated file. Do not modify.
package bootstrap

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"

	"github.com/labbs/castle/config"
)

func InitFiber(logger zerolog.Logger) *fiber.App {
	fconfig := fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	}

	r := fiber.New(fconfig)

	// enable gofiber logs (custom middleware)
	if config.EnableHTTPLogs {
		r.Use(fiberLogger(logger))
	}

	// gofiber recover => https://docs.gofiber.io/api/middleware/recover
	r.Use(recover.New())

	r.Use(cors.New())

	r.Use(compress.New())

	r.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	return r
}

func fiberLogger(logger zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		timeStart := time.Now()
		err := c.Next()
		logger.Info().
			Int("status", c.Response().StatusCode()).
			Dur("duration", time.Since(timeStart)).
			Str("method", string(c.Request().Header.Method())).
			Str("remote_addr", c.IP()).
			Str("path", c.Path()).
			Str("user_agent", c.Get("User-Agent")).
			Int("bytes_sent", c.Response().Header.ContentLength()).
			Int("bytes_received", c.Request().Header.ContentLength()).
			Str("referer", c.Get("Referer")).
			Str("proto", c.Protocol()).
			Str("host", c.Hostname()).
			Msg("")
		return err
	}
}
