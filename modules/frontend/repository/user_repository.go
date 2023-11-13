package repository

import (
	"github.com/goccy/go-json"

	initBootstrap "github.com/labbs/castle/bootstrap"
	pb "github.com/labbs/castle/gen/user"
)

type UserRepository struct {
	BusMessages chan initBootstrap.Message
}

func NewUserRepository(busMessages chan initBootstrap.Message) UserRepository {
	return UserRepository{BusMessages: busMessages}
}

func (d *UserRepository) GetUserByUsername(email string) (*pb.User, error) {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "user:get_by_email", Data: email, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return &pb.User{}, response.Error
	}

	var user pb.User
	err := json.Unmarshal(response.Data.([]byte), &user)
	if err != nil {
		return &pb.User{}, err
	}

	return &user, nil
}
