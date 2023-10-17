package controller

import (
	"github.com/labbs/castle/modules/project/repository"
	"github.com/rs/zerolog"
)

type EnvironmentController struct {
	Repository repository.EnvironmentRepository
	Logger     zerolog.Logger
}
