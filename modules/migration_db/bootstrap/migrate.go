package bootstrap

import (
	"context"
	"embed"

	"github.com/labbs/castle/config"
	"github.com/labbs/castle/internal/zerolog"
	_ "github.com/labbs/castle/modules/migration_db/bootstrap/migrations"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql migrations/*.go
var embedMigrations embed.FS

func MigrateDatabase(app Application, c config.Config) error {
	goose.SetBaseFS(embedMigrations)
	logger := app.Logger.With().Str("event", "migration").Logger()
	goose.SetLogger(&zerolog.ZerologGooseAdapter{Logger: logger})
	if err := goose.SetDialect(c.Database.Engine); err != nil {
		return err
	}

	sqlDB, err := app.Db.DB()
	if err != nil {
		app.Logger.Fatal().Err(err).Msg("failed to get database connection")
	}

	if err := goose.UpContext(context.Background(), sqlDB, "migrations"); err != nil {
		if err.Error() != "no change" {
			app.Logger.Fatal().Err(err).Msg("failed to migrate database")
		}
	}

	return nil
}
