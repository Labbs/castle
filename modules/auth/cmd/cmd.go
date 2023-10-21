package cmd

import (
	initBootstrap "github.com/labbs/castle/bootstrap"

	"github.com/labbs/castle/modules/auth/api/route"
	"github.com/labbs/castle/modules/auth/bootstrap"
)

func Init(initBootstrapApp *initBootstrap.Application) {
	app := bootstrap.App(initBootstrapApp)

	route.Setup(app)
}
