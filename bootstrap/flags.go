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
			Destination: &config.AppConfig.ConfigFile,
		},
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"d"},
			EnvVars:     []string{"DEBUG"},
			Value:       false,
			Usage:       "Enable debug mode",
			Destination: &config.AppConfig.Debug,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "pretty-logs",
			Aliases:     []string{"pl"},
			EnvVars:     []string{"PRETTY_LOGS"},
			Value:       false,
			Usage:       "Enable pretty logs",
			Destination: &config.AppConfig.PrettyLogs,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "local-dev",
			Aliases:     []string{"ld"},
			EnvVars:     []string{"LOCAL_DEV"},
			Value:       false,
			Usage:       "Enable local development mode",
			Destination: &config.AppConfig.LocalDev,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "init-frontend-view-engine",
			Aliases:     []string{"ifve"},
			EnvVars:     []string{"INIT_FRONTEND_VIEW_ENGINE"},
			Value:       true,
			Usage:       "Initialize frontend view engine",
			Destination: &config.AppConfig.InitFrontendViewEngine,
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
			Destination: &config.AppConfig.Database.DSN,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "database.engine",
			Usage:       "Database Engine",
			EnvVars:     []string{"DATABASE_ENGINE"},
			Value:       "sqlite3",
			Destination: &config.AppConfig.Database.Engine,
		}),
	}
}

func SessionFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "session.secret-key",
			Aliases:     []string{"ssk"},
			EnvVars:     []string{"SESSION_SECRET_KEY"},
			Usage:       "Session secret key",
			Destination: &config.AppConfig.Session.SecretKey,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "session.expire",
			Aliases:     []string{"se"},
			EnvVars:     []string{"SESSION_EXPIRE"},
			Usage:       "Session expire time in minutes",
			Destination: &config.AppConfig.Session.Expire,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "session.issuer",
			Aliases:     []string{"si"},
			EnvVars:     []string{"SESSION_ISSUER"},
			Usage:       "Session issuer",
			Destination: &config.AppConfig.Session.Issuer,
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
			Destination: &config.AppConfig.Port,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "enable-http-logs",
			Aliases:     []string{"ehl"},
			EnvVars:     []string{"ENABLE_HTTP_LOGS"},
			Usage:       "Enable http server logs",
			Destination: &config.AppConfig.EnableHTTPLogs,
		}),
	}
	flags = append(flags, GenericFlags()...)
	flags = append(flags, DatabaseFlags()...)
	flags = append(flags, SessionFlags()...)
	return flags
}
