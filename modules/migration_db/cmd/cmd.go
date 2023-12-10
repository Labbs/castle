package cmd

import (
	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/modules/migration_db/bootstrap"
)

func Init(initBootstrapApp *initBootstrap.Application) error {
	return bootstrap.App(initBootstrapApp)
}
