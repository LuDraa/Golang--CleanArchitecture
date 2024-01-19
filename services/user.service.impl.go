package services

import (
	"ecommerce/gmr/interfaces"
	"ecommerce/gmr/models"
)

type UserServiceImpl struct {
	UserDataLayer interfaces.UserDataLayer
}

func NewUserService(UserDataLayer interfaces.UserDataLayer) interfaces.UserService {
	return &UserServiceImpl{
		UserDataLayer: UserDataLayer,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	err := u.UserDataLayer.CreateUser(user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	user, err := u.UserDataLayer.GetUser(name)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	users, err := u.UserDataLayer.GetAll()
	return users, err
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	err := u.UserDataLayer.UpdateUser(user)
	return err
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	err := u.UserDataLayer.DeleteUser(name)
	return err
}
