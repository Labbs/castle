package bootstrap

import (
	"github.com/labbs/castle/internal"
	"github.com/labbs/castle/modules/user/domain"
	"golang.org/x/crypto/bcrypt"
)

func InitOrMigrateDatabase(app Application) error {
	db := app.Db
	logger := app.Logger
	err := db.AutoMigrate(
		&domain.User{},
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error on migrate database")
	}

	var userCount int64
	db.Model(&domain.User{}).Count(&userCount)
	if userCount != 0 {
		return nil
	}

	pwd := internal.RandomString(30)

	bytes, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		logger.Error().Err(err).Msg("Error on generate password for default user")
		return err
	}

	defaultUser := domain.User{
		Username: "default",
		Password: string(bytes),
	}

	err = db.Create(&defaultUser).Error
	if err != nil {
		logger.Error().Err(err).Msg("Error on create default user")
		return err
	}

	logger.Info().Str("username", defaultUser.Username).Str("password", string(pwd)).Msg("Default user created")
	return nil
}
