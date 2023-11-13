package repository

import (
	pb "github.com/labbs/castle/gen/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return UserRepository{database: database}
}

func (d *UserRepository) GetUserByUsername(username string) (*pb.User, error) {
	u := pb.User{}
	r := d.database.Where("username = ?", username).First(&u)
	return &u, r.Error
}

func (d *UserRepository) UpdateUser(user *pb.User) error {
	r := d.database.Save(&user)
	return r.Error
}
