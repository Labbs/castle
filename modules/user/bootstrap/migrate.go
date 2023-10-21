package bootstrap

import (
	"github.com/labbs/castle/config"
	"github.com/labbs/castle/internal"
	"github.com/labbs/castle/modules/user/domain"
	"golang.org/x/crypto/bcrypt"
)

func InitOrMigrateDatabase(app Application, c config.Config) error {
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

	var r []byte
	if !c.LocalDev {
		r = internal.RandomString(30)
	} else {
		r = []byte("default")
	}

	pwd, err := bcrypt.GenerateFromPassword(r, 14)
	if err != nil {
		logger.Error().Err(err).Msg("Error on generate password for default user")
		return err
	}

	defaultUser := domain.User{
		Username: "default",
		Password: string(pwd),
	}

	err = db.Create(&defaultUser).Error
	if err != nil {
		logger.Error().Err(err).Msg("Error on create default user")
		return err
	}

	logger.Info().Str("username", defaultUser.Username).Str("password", string(r)).Msg("Default user created")
	return nil
}
