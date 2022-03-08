package service

import (
	"task/api/repository"
	"task/models"
)

// UserService struct
type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (u UserService) CreateUser(user models.UserRegister) error {
	var newUser models.User
	newUser.Name = user.Name
	newUser.Email = user.Email
	newUser.Password = user.Password
	newUser.Surname = user.Surname
	return u.repo.CreateUser(&newUser)
}

func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)

}
