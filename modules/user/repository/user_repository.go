package repository

import (
	"github.com/labbs/castle/modules/user/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *userRepository {
	return &userRepository{database: database}
}

func (d *userRepository) GetUserByEmail(email string) (domain.User, error) {
	u := domain.User{}
	r := d.database.Where("email = ?", email).First(&u)
	return u, r.Error
}

func (d *userRepository) UpdateUser(user domain.User) error {
	r := d.database.Save(&user)
	return r.Error
}
