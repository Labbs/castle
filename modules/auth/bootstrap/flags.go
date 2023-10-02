package bootstrap

import (
	"github.com/labbs/castle/modules/auth/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func SessionFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "session.secret-key",
			Aliases:     []string{"ssk"},
			EnvVars:     []string{"SESSION_SECRET_KEY"},
			Usage:       "Session secret key",
			Destination: &config.Session.SecretKey,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "session.expire",
			Aliases:     []string{"se"},
			EnvVars:     []string{"SESSION_EXPIRE"},
			Usage:       "Session expire time in minutes",
			Destination: &config.Session.Expire,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "session.issuer",
			Aliases:     []string{"si"},
			EnvVars:     []string{"SESSION_ISSUER"},
			Usage:       "Session issuer",
			Destination: &config.Session.Issuer,
		}),
	}
}
