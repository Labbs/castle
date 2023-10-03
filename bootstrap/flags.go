package bootstrap

import (
	"github.com/labbs/castle/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func GenericFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			EnvVars:     []string{"CONFIG"},
			Usage:       "Config file path",
			Value:       "config.json",
			Destination: &config.ConfigFile,
		},
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"d"},
			EnvVars:     []string{"DEBUG"},
			Value:       false,
			Usage:       "Enable debug mode",
			Destination: &config.Debug,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "pretty-logs",
			Aliases:     []string{"pl"},
			EnvVars:     []string{"PRETTY_LOGS"},
			Value:       false,
			Usage:       "Enable pretty logs",
			Destination: &config.PrettyLogs,
		}),
	}
}

func DatabaseFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "database.dsn",
			Usage:       "Database DSN",
			EnvVars:     []string{"DATABASE_DSN"},
			Value:       "castle.db",
			Destination: &config.Database.DSN,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "database.engine",
			Usage:       "Database Engine",
			EnvVars:     []string{"DATABASE_ENGINE"},
			Value:       "sqlite3",
			Destination: &config.Database.Engine,
		}),
	}
}

func ServerFlags() []cli.Flag {
	flags := []cli.Flag{
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			EnvVars:     []string{"PORT"},
			Usage:       "Server Port",
			Destination: &config.Port,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "enable-http-logs",
			Aliases:     []string{"ehl"},
			EnvVars:     []string{"ENABLE_HTTP_LOGS"},
			Usage:       "Enable http server logs",
			Destination: &config.EnableHTTPLogs,
		}),
	}
	flags = append(flags, GenericFlags()...)
	flags = append(flags, DatabaseFlags()...)
	return flags
}
