package migrations

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/pressly/goose/v3"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	goose.AddMigrationContext(upCreateDefaultUser, downCreateDefaultUser)
}

func upCreateDefaultUser(ctx context.Context, tx *sql.Tx) error {
	r := []byte("default")
	pwd, err := bcrypt.GenerateFromPassword(r, 14)
	if err != nil {
		return err
	}

	id := utils.UUIDv4()

	query := "INSERT INTO user (id, email, password) VALUES ($1, $2, $3)"

	_, err = tx.ExecContext(ctx, query, id, "default@castle.local", pwd)
	if err != nil {
		return err
	}

	return nil
}

func downCreateDefaultUser(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM user WHERE email = 'default@castle.local'")
	return err
}
