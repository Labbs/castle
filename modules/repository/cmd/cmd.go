package cmd

import (
	initBootstrap "github.com/labbs/castle/bootstrap"

	"github.com/labbs/castle/modules/repository/api/route"
	"github.com/labbs/castle/modules/repository/bootstrap"
	"github.com/labbs/castle/modules/repository/bus"
)

func Init(initBootstrapApp *initBootstrap.Application) {
	app := bootstrap.App(initBootstrapApp)

	bus.Setup(app)
	route.Setup(app)
}
