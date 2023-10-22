package cmd

import (
	initBootstrap "github.com/labbs/castle/bootstrap"

	"github.com/labbs/castle/modules/scheduler/api/route"
	"github.com/labbs/castle/modules/scheduler/bootstrap"
	"github.com/labbs/castle/modules/scheduler/bus"
)

func Init(initBootstrapApp *initBootstrap.Application) {
	app := bootstrap.App(initBootstrapApp)

	bus.Setup(app)
	route.Setup(app)
}
