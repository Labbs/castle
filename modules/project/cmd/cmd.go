package cmd

import (
	initBootstrap "github.com/labbs/castle/bootstrap"

	"github.com/labbs/castle/modules/project/api/route"
	"github.com/labbs/castle/modules/project/bootstrap"
	"github.com/labbs/castle/modules/project/bus"
)

func Init(initBootstrapApp *initBootstrap.Application) {
	app := bootstrap.App(initBootstrapApp)

	bus.Setup(app)
	route.Setup(app)
}
