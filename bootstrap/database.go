package bootstrap

import (
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"

	"github.com/labbs/castle/config"
)

func InitDatabase(logger zerolog.Logger, c config.Config) *gorm.DB {
	newLogger := l.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		l.Config{
			SlowThreshold:             time.Minute, // Slow SQL threshold
			LogLevel:                  l.Silent,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	var db *gorm.DB
	var err error

	switch engine := config.AppConfig.Database.Engine; engine {
	case "mysql":
		db, err = gorm.Open(mysql.Open(c.Database.DSN), &gorm.Config{Logger: newLogger})
	case "postgres":
		db, err = gorm.Open(postgres.Open(c.Database.DSN), &gorm.Config{Logger: newLogger})
	default:
		db, err = gorm.Open(sqlite.Open(c.Database.DSN), &gorm.Config{Logger: newLogger})
	}

	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}

	return db
}
