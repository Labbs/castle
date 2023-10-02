// Path: cmd/main.go
// This is a generated file. Do not modify.
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/config"

	authBootstrap "github.com/labbs/castle/modules/auth/bootstrap"
	authModule "github.com/labbs/castle/modules/auth/cmd"

	coreModule "github.com/labbs/castle/modules/core/cmd"
	userModule "github.com/labbs/castle/modules/user/cmd"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

// Version is the current version of the application.
// This value is overwritten by the build process.
var version = "development"
var commands []*cli.Command = []*cli.Command{}

func main() {
	flags := bootstrap.ServerFlags()
	flags = append(flags, authBootstrap.SessionFlags()...)

	// Initialize the cli
	app := cli.NewApp()
	app.Name = "Castle"
	app.Version = version

	app.Commands = []*cli.Command{
		{
			Name:   "server",
			Usage:  "Start the monolith server",
			Flags:  flags,
			Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewJSONSourceFromFlagFunc("config")),
			Action: func(c *cli.Context) error {
				appBootstrap := bootstrap.App()

				if config.Debug {
					appBootstrap.Logger.Debug().Interface("fiber.routes", appBootstrap.Fiber.GetRoutes()).Msg("fiber routes")
				}

				authModule.Init(appBootstrap)
				userModule.Init(appBootstrap)
				coreModule.Init(appBootstrap)

				appBootstrap.Logger.Info().Str("event", "server.start").Msg("app listening on 0.0.0.0:" + strconv.Itoa(config.Port))
				return appBootstrap.Fiber.Listen(":" + strconv.Itoa(config.Port))
			},
		},
	}

	app.Commands = append(app.Commands, commands...)

	config.Version = version
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
