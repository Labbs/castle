package service

import "github.com/labbs/castle/modules/user/domain"

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) GetUserByEmail(email string) (domain.User, error) {
	return s.userRepository.GetUserByEmail(email)
}

func (s *userService) UpdateUser(user domain.User) error {
	return s.userRepository.UpdateUser(user)
}
