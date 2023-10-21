package repository

import (
	"github.com/labbs/castle/modules/user/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return UserRepository{database: database}
}

func (d *UserRepository) GetUserByUsername(username string) (domain.User, error) {
	u := domain.User{}
	r := d.database.Where("username = ?", username).First(&u)
	return u, r.Error
}

func (d *UserRepository) UpdateUser(user domain.User) error {
	r := d.database.Save(&user)
	return r.Error
}
