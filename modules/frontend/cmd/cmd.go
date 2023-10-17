package cmd

import (
	initBootstrap "github.com/labbs/castle/bootstrap"

	"github.com/labbs/castle/modules/frontend/app/route"
	"github.com/labbs/castle/modules/frontend/bootstrap"
)

func Init(initBootstrapApp *initBootstrap.Application) {
	app := bootstrap.App(initBootstrapApp)

	route.Setup(app)
}
