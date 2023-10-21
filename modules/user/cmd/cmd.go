package cmd

import (
	initBootstrap "github.com/labbs/castle/bootstrap"

	"github.com/labbs/castle/modules/user/api/route"
	"github.com/labbs/castle/modules/user/bootstrap"
	"github.com/labbs/castle/modules/user/bus"
)

func Init(initBootstrapApp *initBootstrap.Application) {
	app := bootstrap.App(initBootstrapApp)

	bus.Setup(app)
	route.Setup(app)
}
