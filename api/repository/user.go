package repository

import (
	"task/infrastructure"
	"task/models"
	"task/util"
)

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (u UserRepository) CreateUser(user *models.User) error {
	return u.db.DB.Create(user).Error
}

func (u UserRepository) LoginUser(user models.UserLogin) (*models.User, error) {

	var dbUser models.User
	email := user.Email
	password := user.Password

	err := u.db.DB.Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	hashErr := util.CheckPassword(password, dbUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	return &dbUser, nil
}
