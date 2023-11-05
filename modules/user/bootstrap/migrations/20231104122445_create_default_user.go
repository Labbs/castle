package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/labbs/castle/config"
	"github.com/labbs/castle/internal"
	"github.com/pressly/goose/v3"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	goose.AddMigrationContext(upCreateDefaultUser, downCreateDefaultUser)
}

func upCreateDefaultUser(ctx context.Context, tx *sql.Tx) error {
	var r []byte
	if !config.AppConfig.LocalDev {
		r = internal.RandomString(30)
	} else {
		r = []byte("default")
	}

	pwd, err := bcrypt.GenerateFromPassword(r, 14)
	if err != nil {
		return err
	}

	id := utils.UUIDv4()

	_, err = tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO user (id, email, password) VALUES (%s, %s, %s)", id, "default@castle.local", pwd))
	if err != nil {
		return err
	}
	return nil
}

func downCreateDefaultUser(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM user WHERE email = 'default@castle.local'")
	return err
}
