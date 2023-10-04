package repository

import (
	initBootstrap "github.com/labbs/castle/bootstrap"
)

type UserRepository struct {
	BusMessages chan initBootstrap.Message
}

func NewUserRepository(busMessages chan initBootstrap.Message) UserRepository {
	return UserRepository{BusMessages: busMessages}
}
