package cmd

import (
	initBootstrap "github.com/labbs/castle/bootstrap"

	"github.com/labbs/castle/modules/frontend/bootstrap"
)

func Init(initBootstrapApp *initBootstrap.Application) {
	bootstrap.App(initBootstrapApp)
}
