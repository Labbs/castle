package cmd

import (
	initBootstrap "github.com/labbs/castle/bootstrap"

	"github.com/labbs/castle/modules/task/api/route"
	"github.com/labbs/castle/modules/task/bootstrap"
	"github.com/labbs/castle/modules/task/bus"
)

func Init(initBootstrapApp *initBootstrap.Application) {
	app := bootstrap.App(initBootstrapApp)

	bus.Setup(app)
	route.Setup(app)
}
