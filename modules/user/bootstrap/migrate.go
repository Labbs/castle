package bootstrap

import (
	"context"
	"embed"

	"github.com/labbs/castle/config"
	"github.com/pressly/goose/v3"
	"gorm.io/gorm/logger"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func MigrateDatabase(app Application, c config.Config) error {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect(c.Database.Engine); err != nil {
		return err
	}

	_ctx := context.Background()
	ctx := context.WithValue(_ctx, "logger", app.Logger())

	logger.Info().Str("username", defaultUser.Username).Str("password", string(r)).Msg("Default user created")
	return nil
}
