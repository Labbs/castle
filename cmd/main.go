// Path: cmd/main.go
// This is a generated file. Do not modify.
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/config"

	authModule "github.com/labbs/castle/modules/auth/cmd"
	frontendModule "github.com/labbs/castle/modules/frontend/cmd"
	migrationDbModule "github.com/labbs/castle/modules/migration_db/cmd"
	projectModule "github.com/labbs/castle/modules/project/cmd"
	repositoryModule "github.com/labbs/castle/modules/repository/cmd"
	taskModule "github.com/labbs/castle/modules/task/cmd"

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
				// Init fiber, logger, database, etc
				appBootstrap := bootstrap.App(&config.AppConfig)

				// Send a message to inform that the app is running in local development mode
				if config.AppConfig.LocalDev {
					appBootstrap.Logger.Info().Msg("Running in local development mode")
				}

				err := migrationDbModule.Init(appBootstrap)
				if err != nil {
					appBootstrap.Logger.Fatal().Err(err).Msg("Failed to migrate database")
				}

				// Initialize modules
				authModule.Init(appBootstrap)
				userModule.Init(appBootstrap)
				projectModule.Init(appBootstrap)
				taskModule.Init(appBootstrap)
				repositoryModule.Init(appBootstrap)
				frontendModule.Init(appBootstrap)

				// Start the fiber server
				appBootstrap.Logger.Info().Str("event", "server.start").Msg("app listening on 0.0.0.0:" + strconv.Itoa(config.AppConfig.Port))
				return appBootstrap.Fiber.Listen(":" + strconv.Itoa(config.AppConfig.Port))
			},
		},
	}

	app.Commands = append(app.Commands, commands...)

	config.AppConfig.Version = version
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
